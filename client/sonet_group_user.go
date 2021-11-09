package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) Sonet_groupUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupUserInvite(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.user.invite", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupUserRequest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.user.request", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupUserAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.user.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupUserUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.user.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupUserDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.user.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_groupUserGroups(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.user.groups", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
