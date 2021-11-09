package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MobileUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileUserCanusetelephony(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.user.canusetelephony", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
