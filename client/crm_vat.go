package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmVatFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.vat.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmVatList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.vat.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmVatGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.vat.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmVatAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.vat.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmVatUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.vat.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmVatDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.vat.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
