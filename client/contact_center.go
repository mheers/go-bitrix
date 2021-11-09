package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) _placements(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("_placements", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
