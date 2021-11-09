package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskCtaskelapseditemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskelapseditem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskelapseditemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskelapseditem.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskelapseditemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskelapseditem.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskelapseditemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskelapseditem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskelapseditemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskelapseditem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskelapseditemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskelapseditem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskelapseditemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskelapseditem.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
