package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmStatusFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusEntityTypes(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.entity.types", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusEntityItems(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.entity.items", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmStatusExtraFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.status.extra.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
