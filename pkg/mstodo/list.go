package mstodo

import "net/http"

// Retrieves all task lists
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

// Retrieves a specific list
func (c *Client) GetList(listId string) (List, error) {
	reqUrl := EndpointV1 + "lists/" + listId
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return List{}, err
	}

	var list List
	err = c.Do(req, &list)
	if err != nil {
		return List{}, err
	}
	return list, nil
}
