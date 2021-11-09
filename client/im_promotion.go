package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImPromotionGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.promotion.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImPromotionRead(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.promotion.read", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
