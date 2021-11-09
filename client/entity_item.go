package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) EntityItemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityItemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityItemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityItemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityItemPropertyAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.property.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityItemPropertyGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.property.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityItemPropertyUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.property.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityItemPropertyDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.item.property.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
