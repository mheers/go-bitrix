package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmProductrowFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productrow.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductrowAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productrow.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductrowGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productrow.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductrowList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productrow.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductrowUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productrow.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmProductrowDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.productrow.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
