package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmActivityFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityCommunicationFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.communication.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityTypeAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.type.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityTypeList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.type.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmActivityTypeDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.activity.type.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
