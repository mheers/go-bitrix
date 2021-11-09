package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) EntityAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntityRights(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.rights", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
