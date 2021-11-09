package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmExternalchannelConnectorFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.connector.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmExternalchannelConnectorList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.connector.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmExternalchannelConnectorRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.connector.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmExternalchannelConnectorUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.connector.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmExternalchannelCompany(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.company", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmExternalchannelContact(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.contact", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmExternalchannelActivityCompany(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.activity.company", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmExternalchannelActivityContact(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.externalchannel.activity.contact", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
