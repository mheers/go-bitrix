package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) SalePaysystemHandlerAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.handler.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemHandlerUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.handler.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemHandlerDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.handler.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemHandlerList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.handler.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemSettingsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.settings.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemSettingsUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.settings.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemSettingsInvoiceGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.settings.invoice.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemSettingsPaymentGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.settings.payment.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemPayInvoice(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.pay.invoice", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) SalePaysystemPayPayment(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sale.paysystem.pay.payment", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
