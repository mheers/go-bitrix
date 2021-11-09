package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImopenlinesWidgetConfigGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.config.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesWidgetDialogGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.dialog.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesWidgetUserRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.user.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesWidgetUserConsentApply(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.user.consent.apply", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesWidgetUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesWidgetOperatorGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.operator.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesWidgetVoteSend(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.vote.send", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesWidgetFormSend(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.widget.form.send", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
