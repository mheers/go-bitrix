package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImMessageAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.message.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImMessageDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.message.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImMessageUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.message.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImMessageLike(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.message.like", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImMessageCommand(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.message.command", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImMessageShare(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.message.share", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImMessageUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.message.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
