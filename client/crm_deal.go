package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmDealFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealProductrowsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.productrows.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealProductrowsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.productrows.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealContactFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.contact.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealContactAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.contact.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealContactDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.contact.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealContactItemsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.contact.items.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealContactItemsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.contact.items.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealContactItemsDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.contact.items.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealRecurringFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.recurring.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealRecurringList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.recurring.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealRecurringAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.recurring.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealRecurringGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.recurring.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealRecurringUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.recurring.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealRecurringDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.recurring.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealRecurringExpose(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.recurring.expose", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealUserfieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.userfield.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealDetailsConfigurationGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.details.configuration.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealDetailsConfigurationSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.details.configuration.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealDetailsConfigurationReset(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.details.configuration.reset", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealDetailsConfigurationForcecommonscopeforall(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.deal.details.configuration.forcecommonscopeforall", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
