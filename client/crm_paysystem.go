package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmPaysystemFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.paysystem.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmPaysystemList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.paysystem.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
