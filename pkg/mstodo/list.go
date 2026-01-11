package mstodo

import "net/http"

func (c *Client) GetLists() ([]List, error) {
	reqUrl := EndpointV1 + "lists"
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var lists []List
	err = c.PaginatedDo(req, &lists)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
