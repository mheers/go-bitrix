package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskLogitemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.logitem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskLogitemList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.logitem.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
