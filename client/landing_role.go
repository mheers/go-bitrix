package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingRoleGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.role.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRoleGetrights(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.role.getrights", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRoleSetrights(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.role.setrights", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRoleSetaccesscodes(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.role.setaccesscodes", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRoleIsenabled(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.role.isenabled", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRoleEnable(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.role.enable", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
