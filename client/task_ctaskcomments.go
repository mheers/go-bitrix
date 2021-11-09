package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskCtaskcommentsGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcomments.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskcommentsAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcomments.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
