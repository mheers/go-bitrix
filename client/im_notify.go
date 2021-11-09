package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImNotifyPersonalAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.notify.personal.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImNotifySystemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.notify.system.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImNotifyDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.notify.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImNotifyRead(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.notify.read", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
