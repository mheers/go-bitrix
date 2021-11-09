package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmProductFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyTypes(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.types", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertySettingsFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.settings.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyEnumerationFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.enumeration.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductPropertyDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.product.property.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
