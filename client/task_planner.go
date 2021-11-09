package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskPlannerGetcurrenttaskslist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.planner.getcurrenttaskslist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskPlannerGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.planner.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
