package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmWebformList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.webform.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmWebformResultAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.webform.result.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmWebformConfigurationGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.webform.configuration.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
