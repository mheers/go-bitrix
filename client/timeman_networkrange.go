package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TimemanNetworkrangeGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.networkrange.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanNetworkrangeSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.networkrange.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanNetworkrangeCheck(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.networkrange.check", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
