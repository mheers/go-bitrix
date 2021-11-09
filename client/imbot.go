package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImbotRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotSendtyping(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.sendtyping", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
