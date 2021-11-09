package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) Sonet_groupFeatureAccess(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group.feature.access", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
