package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmCurrencyFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyLocalizationsFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.localizations.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyLocalizationsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.localizations.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyLocalizationsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.localizations.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyLocalizationsDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.localizations.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyBaseSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.base.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCurrencyBaseGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.currency.base.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
