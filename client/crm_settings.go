package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmSettingsModeGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.settings.mode.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
