package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) BizprocTaskList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.task.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocTaskComplete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.task.complete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
