package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImMobileConfigGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.mobile.config.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
