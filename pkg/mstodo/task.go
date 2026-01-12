package mstodo

import (
	"net/http"
)

// Gets all tasks for a given list
// If completed is set to true, it will
// also fetch completed tasks, otherwise it will only retrieve uncompleted task
func (c *Client) GetListTasks(listId string, completed bool) ([]Task, error) {
	reqUrl := EndpointV1 + "lists/" + listId + "/tasks"
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	if !completed {
		q := req.URL.Query()
		q.Add("$filter", "status ne 'completed'")
		req.URL.RawQuery = q.Encode()
	}

	var tasks []Task
	err = c.PaginatedDo(req, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Retrieves all tasks for all lists
// If completed is set to true, will also retrieve completed tasks
func (c *Client) GetAllTasks(completed bool) ([]List, error) {
	lists, err := c.GetLists()
	if err != nil {
		return nil, err
	}
	for i := range lists {
		tasks, err := c.GetListTasks(lists[i].Id, completed)
		if err != nil {
			return nil, err
		}
		lists[i].Tasks = &tasks
	}
	return lists, nil
}
