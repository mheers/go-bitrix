package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImRecentGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.recent.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImRecentList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.recent.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImRecentPin(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.recent.pin", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImRecentHide(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.recent.hide", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImRecentUnread(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.recent.unread", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
