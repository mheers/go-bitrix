package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskCommentGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.comment.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCommentAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.comment.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
