package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) UserHistoryList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.history.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserHistoryFieldsList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("user.history.fields.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
