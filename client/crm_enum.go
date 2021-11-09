package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmEnumSettingsMode(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.settings.mode", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumOwnertype(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.ownertype", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumAddresstype(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.addresstype", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumContenttype(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.contenttype", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumActivitytype(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.activitytype", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumActivitypriority(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.activitypriority", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumActivitydirection(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.activitydirection", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumActivitynotifytype(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.activitynotifytype", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumActivitystatus(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.activitystatus", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmEnumEntityeditorConfigurationScope(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.enum.entityeditor.configuration.scope", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
