package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmMeasureFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.measure.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmMeasureAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.measure.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmMeasureGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.measure.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmMeasureList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.measure.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmMeasureUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.measure.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmMeasureDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.measure.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
