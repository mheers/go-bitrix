package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmInvoiceFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceGetexternallink(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.getexternallink", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceStatusFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.status.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceStatusAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.status.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceStatusGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.status.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceStatusList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.status.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceStatusUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.status.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceStatusDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.status.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceRecurringFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.recurring.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceRecurringList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.recurring.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceRecurringAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.recurring.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceRecurringGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.recurring.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceRecurringUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.recurring.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceRecurringDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.recurring.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceRecurringExpose(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.recurring.expose", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceUserfieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.userfield.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmInvoiceUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.invoice.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
