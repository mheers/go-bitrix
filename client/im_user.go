package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserListGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.list.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserBusinessList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.business.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserBusinessGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.business.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserStatusGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.status.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserStatusSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.status.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserStatusIdleStart(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.status.idle.start", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserStatusIdleEnd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.status.idle.end", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImUserStatusIdleContinue(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.user.status.idle.continue", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
