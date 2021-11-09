package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) AppOptionGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("app.option.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) AppOptionSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("app.option.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) AppInfo(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("app.info", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
