package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImopenlinesBotSessionOperator(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.bot.session.operator", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesBotSessionSendMessage(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.bot.session.send.message", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesBotSessionMessageSend(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.bot.session.message.send", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesBotSessionTransfer(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.bot.session.transfer", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesBotSessionFinish(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.bot.session.finish", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
