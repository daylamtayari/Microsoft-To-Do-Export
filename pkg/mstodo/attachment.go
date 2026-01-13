package mstodo

import "net/http"

// Retrieves all attachments for a given task
// NOTE: The ContentBytes field of the attachment objects is not
// populated on a listing of attachments, it is only populated when
// retrieving the attachment directly with GetAttachment
func (c *Client) ListAttachments(listId string, taskId string) ([]Attachment, error) {
	reqUrl := EndpointV1 + "lists/" + listId + "/tasks/" + taskId + "/attachments"
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var attachments []Attachment
	err = c.PaginatedDo(req, &attachments)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

// Retrieves a given attachment, including its contents that are stored in the
// ContentBytes field in a base64 encoded way
func (c *Client) GetAttachment(listId string, taskId string, attachmentId string) (Attachment, error) {
	reqUrl := EndpointV1 + "lists/" + listId + "/tasks/" + taskId + "/attachments/" + attachmentId
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return Attachment{}, err
	}

	var attachment Attachment
	err = c.Do(req, &attachment)
	if err != nil {
		return Attachment{}, err
	}
	return attachment, nil
}
