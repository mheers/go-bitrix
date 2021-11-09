package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingTemplateGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.template.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingTemplateSetsiteref(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.template.setsiteref", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingTemplateSetlandingref(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.template.setlandingref", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingTemplateGetsiteref(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.template.getsiteref", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingTemplateGetlandingref(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.template.getlandingref", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
