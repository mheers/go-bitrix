package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImbotAppRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.app.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotAppUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.app.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotAppUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.app.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
