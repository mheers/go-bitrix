package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImbotCommandRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.command.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotCommandUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.command.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotCommandUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.command.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotCommandAnswer(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.command.answer", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
