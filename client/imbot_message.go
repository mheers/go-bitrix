package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImbotMessageAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.message.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotMessageDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.message.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotMessageUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.message.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotMessageLike(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.message.like", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
