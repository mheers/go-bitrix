package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskElapseditemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.elapseditem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskElapseditemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.elapseditem.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskElapseditemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.elapseditem.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskElapseditemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.elapseditem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskElapseditemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.elapseditem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskElapseditemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.elapseditem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskElapseditemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.elapseditem.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
