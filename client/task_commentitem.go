package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskCommentitemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.commentitem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCommentitemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.commentitem.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCommentitemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.commentitem.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCommentitemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.commentitem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCommentitemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.commentitem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCommentitemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.commentitem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCommentitemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.commentitem.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
