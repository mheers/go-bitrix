package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MobilePushTypesGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.push.types.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobilePushConfigGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.push.config.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobilePushConfigSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.push.config.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobilePushStatusGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.push.status.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobilePushStatusSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.push.status.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobilePushSmartfilterStatusGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.push.smartfilter.status.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobilePushSmartfilterStatusSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.push.smartfilter.status.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
