package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmCatalogFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.catalog.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCatalogGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.catalog.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmCatalogList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.catalog.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
