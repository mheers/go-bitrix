package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmLeadFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadProductrowsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.productrows.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadProductrowsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.productrows.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadUserfieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.userfield.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadDetailsConfigurationGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.details.configuration.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadDetailsConfigurationSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.details.configuration.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadDetailsConfigurationReset(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.details.configuration.reset", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmLeadDetailsConfigurationForcecommonscopeforall(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.lead.details.configuration.forcecommonscopeforall", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
