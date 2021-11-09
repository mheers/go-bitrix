package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingSyspageSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.syspage.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSyspageGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.syspage.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSyspageDeleteforsite(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.syspage.deleteforsite", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSyspageDeleteforlanding(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.syspage.deleteforlanding", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSyspageGetspecialpage(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.syspage.getspecialpage", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
