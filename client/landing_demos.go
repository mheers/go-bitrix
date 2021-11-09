package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingDemosGetsitelist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.demos.getsitelist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingDemosGetpagelist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.demos.getpagelist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingDemosGeturlpreview(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.demos.geturlpreview", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingDemosRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.demos.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingDemosUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.demos.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingDemosGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.demos.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
