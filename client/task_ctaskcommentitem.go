package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskCtaskcommentitemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcommentitem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskcommentitemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcommentitem.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskcommentitemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcommentitem.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskcommentitemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcommentitem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskcommentitemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcommentitem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskcommentitemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcommentitem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskcommentitemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskcommentitem.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
