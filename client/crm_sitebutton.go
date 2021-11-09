package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmSitebuttonConfigurationGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.sitebutton.configuration.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
