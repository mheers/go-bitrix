package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) Sonet_groupGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupCreate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.create", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupSetowner(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.setowner", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
