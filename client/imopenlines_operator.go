package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImopenlinesOperatorAnswer(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.operator.answer", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesOperatorSkip(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.operator.skip", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesOperatorSpam(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.operator.spam", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesOperatorTransfer(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.operator.transfer", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesOperatorFinish(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.operator.finish", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
