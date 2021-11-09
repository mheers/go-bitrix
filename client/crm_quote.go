package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmQuoteFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteProductrowsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.productrows.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteProductrowsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.productrows.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteContactFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.contact.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteContactAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.contact.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteContactDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.contact.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteContactItemsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.contact.items.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteContactItemsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.contact.items.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteContactItemsDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.contact.items.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteUserfieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.userfield.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmQuoteUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.quote.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
