package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) DiskRightsGettasks(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.rights.gettasks", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
