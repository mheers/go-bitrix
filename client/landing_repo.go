package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingRepoCheckcontent(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.checkcontent", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRepoRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRepoUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRepoGetappinfo(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.getappinfo", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRepoBind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.bind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRepoUnbind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.unbind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRepoGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
