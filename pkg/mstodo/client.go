package mstodo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

const EndpointV1 = "https://graph.microsoft.com/v1.0/me/todo/"

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
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	var body []byte
	if res.ContentLength != 0 && (res.StatusCode == 200 || res.StatusCode == 201) {
		defer res.Body.Close() //nolint:errcheck
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

// A wrapper of the Do method for requests with paginated responses.
// Requires v to be a list type
func (c *Client) PaginatedDo(req *http.Request, v any) error {
	// Have to read and then re-assign the request body value since it can only be read once.
	var reqBody []byte
	var err error
	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		req.Body = io.NopCloser(bytes.NewReader(body))
	} else {
		reqBody = nil
	}
	reqCtx := req.Context()
	reqValues := req.URL.Query()

	var res []byte
	err = c.Do(req, &res)
	if err != nil {
		return err
	}

	nextLink := jsoniter.Get(res, "@odata.nextLink")
	values := jsoniter.Get(res, "value")
	if values == nil || values.ValueType() == jsoniter.InvalidValue {
		return ErrValueRetrieval
	}

	if nextLink == nil || nextLink.ValueType() == jsoniter.InvalidValue {
		err = jsoniter.UnmarshalFromString(values.ToString(), &v)
		return err
	} else {
		combinedValues := values.ToString()
		// Remove final square bracket to allow for a single JSON array
		combinedValues = combinedValues[:len(combinedValues)-1]
		hasNextLink := true
		for hasNextLink {
			skipValue := getSkipValue(nextLink.ToString())
			if skipValue == -1 {
				hasNextLink = false
				continue
			}
			nextReq := req.Clone(reqCtx)
			nextReq.Body = io.NopCloser(bytes.NewReader(reqBody))
			reqValues.Set("$skip", strconv.Itoa(skipValue))
			nextReq.URL.RawQuery = reqValues.Encode()
			res = []byte{}
			err = c.Do(nextReq, &res)
			if err != nil {
				return err
			}

			values := jsoniter.Get(res, "value")
			if values == nil || values.ValueType() == jsoniter.InvalidValue {
				return ErrValueRetrieval
			}
			reqValues := values.ToString()
			// Remove leading and trailing square backets as well as add a leading comma
			reqValues = reqValues[1 : len(reqValues)-1]
			combinedValues += "," + reqValues

			nextLink = jsoniter.Get(res, "@odata.nextLink")
			if nextLink == nil || nextLink.ValueType() == jsoniter.InvalidValue {
				hasNextLink = false
			}
		}
		combinedValues += "]"
		err = jsoniter.UnmarshalFromString(combinedValues, &v)
		if err != nil {
			return err
		}
	}
	return nil
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
