package mstodo

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

const EndpointV1 = "https://graph.microsoft.com/v1.0/me/todo/"

// maxRetries is the number of additional attempts made after a request
// fails with a retryable status code (429 or 5xx).
const maxRetries = 4

var nextLinkRe = regexp.MustCompile(`\$skip=([0-9]+)`)

var (
	ErrValueRetrieval   = errors.New("failed to retrieve the value of the values JSON field")
	ErrUnauthorized     = errors.New("authentication information is either missing or invalid")
	ErrForbidden        = errors.New("access is forbidden to the requested resource")
	ErrNotFound         = errors.New("requested resource does not exist")
	ErrNonHandledStatus = errors.New("non-handled status code returned")
)

// Client manages communication with the Microsoft graph API
type Client struct {
	client *http.Client
}

// Creates a new mstodo API client. If a nil httpClient is provided,
// http.DefaultClient will be used. To use API methods which require
// authentication, provide a token string value or provide an http
// client that will perform the authentication for you and provide nil
// for the token.
func NewClient(httpClient *http.Client, token *string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if token != nil {
		authToken := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: *token,
		})
		httpClient = oauth2.NewClient(context.Background(), authToken)
	}
	return &Client{
		client: httpClient,
	}
}

// Do performs an API request, handles the response,
// and unmarshals the response into a given interface.
// The value to unmarshal must be a pointer to an interface.
// If a pointer to a byte array is provided, the value will be
// the value of the body.
func (c *Client) Do(req *http.Request, v any) error {
	res, err := c.doWithRetry(req)
	if err != nil {
		return err
	}
	defer res.Body.Close() //nolint:errcheck

	var body []byte
	if res.ContentLength != 0 && (res.StatusCode == 200 || res.StatusCode == 201) {
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if _, ok := v.(*[]byte); ok {
			// If a byte array was provided, the body will be
			// returned directly and not unmarshalled.
			*v.(*[]byte) = body
		} else if values := jsoniter.Get(body, "value"); values == nil || values.ValueType() == jsoniter.InvalidValue {
			err = jsoniter.Unmarshal(body, &v)
		} else {
			err = jsoniter.UnmarshalFromString(values.ToString(), &v)
		}
		if err != nil {
			return err
		}
	}
	switch res.StatusCode {
	case 200:
		return nil
	case 201:
		return nil
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	default:
		return fmt.Errorf("%w: %d", ErrNonHandledStatus, res.StatusCode)
	}
}

// doWithRetry sends req, retrying on throttling (429) and transient server
// errors (5xx) up to maxRetries times. It honors the Retry-After header
// when present, falling back to exponential backoff otherwise. All Graph
// requests are bodyless GETs, so req is safe to resend as-is.
func (c *Client) doWithRetry(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		res, err := c.client.Do(req)
		if err != nil {
			return nil, err
		}
		if attempt == maxRetries || !isRetryableStatus(res.StatusCode) {
			return res, nil
		}
		delay := retryDelay(res.Header.Get("Retry-After"), attempt)
		res.Body.Close() //nolint:errcheck
		time.Sleep(delay)
	}
}

// isRetryableStatus reports whether a response status code represents a
// transient failure worth retrying: throttling (429) or a server error.
func isRetryableStatus(statusCode int) bool {
	return statusCode == http.StatusTooManyRequests || statusCode >= 500
}

// retryDelay determines how long to wait before the next retry attempt,
// preferring the server-provided Retry-After value (in seconds) and
// falling back to exponential backoff (1s, 2s, 4s, 8s, ...).
func retryDelay(retryAfterHeader string, attempt int) time.Duration {
	if retryAfterHeader != "" {
		if seconds, err := strconv.ParseFloat(retryAfterHeader, 64); err == nil {
			return time.Duration(seconds * float64(time.Second))
		}
	}
	return time.Duration(1<<attempt) * time.Second
}

// A wrapper of the Do method for requests with paginated responses.
// Requires v to be a list type. The request is expected to be bodyless
// (every paginated Graph endpoint used here is a GET).
func (c *Client) PaginatedDo(req *http.Request, v any) error {
	reqCtx := req.Context()
	reqValues := req.URL.Query()

	var res []byte
	err := c.Do(req, &res)
	if err != nil {
		return err
	}

	// Retrieve all of the elements across multiple pages of data
	var elements []string
	for {
		values := jsoniter.Get(res, "value")
		if values == nil || values.ValueType() == jsoniter.InvalidValue {
			return ErrValueRetrieval
		}
		// Strip the enclosing square brackets of the page's array and check if not empty
		page := values.ToString()
		if inner := page[1 : len(page)-1]; inner != "" {
			elements = append(elements, inner)
		}

		nextLink := jsoniter.Get(res, "@odata.nextLink")
		if nextLink == nil || nextLink.ValueType() == jsoniter.InvalidValue {
			break
		}
		skipValue := getSkipValue(nextLink.ToString())
		if skipValue == -1 {
			break
		}

		nextReq := req.Clone(reqCtx)
		reqValues.Set("$skip", strconv.Itoa(skipValue))
		nextReq.URL.RawQuery = reqValues.Encode()
		res = []byte{}
		err = c.Do(nextReq, &res)
		if err != nil {
			return err
		}
	}

	combinedValues := "[" + strings.Join(elements, ",") + "]"
	return jsoniter.UnmarshalFromString(combinedValues, &v)
}

// Gets the skip value from a next link
// Returns the skip value or -1 if not found
func getSkipValue(nextLink string) int {
	matches := nextLinkRe.FindStringSubmatch(nextLink)
	if len(matches) == 0 {
		return -1
	}
	skipValue, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1
	}
	return skipValue
}
