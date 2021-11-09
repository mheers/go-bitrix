package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmProductsectionFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productsection.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductsectionAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productsection.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductsectionGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productsection.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductsectionList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productsection.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductsectionUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productsection.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductsectionDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productsection.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
