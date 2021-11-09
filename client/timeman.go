package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TimemanSettings(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.settings", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanStatus(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.status", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanOpen(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.open", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanClose(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.close", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanPause(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.pause", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
