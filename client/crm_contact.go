package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmContactFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactCompanyFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.company.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactCompanyAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.company.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactCompanyDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.company.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactCompanyItemsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.company.items.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactCompanyItemsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.company.items.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactCompanyItemsDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.company.items.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactUserfieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.userfield.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactDetailsConfigurationGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.details.configuration.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactDetailsConfigurationSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.details.configuration.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactDetailsConfigurationReset(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.details.configuration.reset", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmContactDetailsConfigurationForcecommonscopeforall(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.contact.details.configuration.forcecommonscopeforall", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
