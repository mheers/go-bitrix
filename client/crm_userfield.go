package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmUserfieldFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.userfield.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmUserfieldTypes(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.userfield.types", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmUserfieldEnumerationFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.userfield.enumeration.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmUserfieldSettingsFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.userfield.settings.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
