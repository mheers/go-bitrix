package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MobileCounterTypesGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.counter.types.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileCounterGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.counter.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileCounterConfigGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.counter.config.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileCounterConfigSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.counter.config.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
