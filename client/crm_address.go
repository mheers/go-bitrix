package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmAddressFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.address.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAddressAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.address.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAddressUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.address.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAddressList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.address.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAddressDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.address.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
