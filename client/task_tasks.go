package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TasksTaskList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("tasks.task.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TasksTaskGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("tasks.task.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TasksTaskAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("tasks.task.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	result := resp.Result()
	return result.(*types.Response), err
}

func (c *Client) TasksTaskUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("tasks.task.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TasksTaskDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("tasks.task.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
