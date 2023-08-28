package client

import (
	"github.com/nightwriter/go-bitrix/types"
	"github.com/nightwriter/go-bitrix/types/users"
)

func (c *Client) UserOptionGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.option.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserOptionSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.option.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserAdmin(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.admin", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserAccess(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.access", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserCurrent(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.current", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserGet(data interface{}) (*users.UsersResponse, error) {
	resp, err := c.DoRaw("user.get", data, &users.UsersResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*users.UsersResponse), err
}

func (c *Client) UserSearch(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.search", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserOnline(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.online", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserCounters(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.counters", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
