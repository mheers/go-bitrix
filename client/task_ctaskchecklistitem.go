package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskCtaskchecklistitemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemComplete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.complete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemRenew(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.renew", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemMoveafteritem(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.moveafteritem", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskchecklistitemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskchecklistitem.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
