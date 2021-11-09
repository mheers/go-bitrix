package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmButtonList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.button.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmButtonWidgetsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.button.widgets.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmButtonGuestRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.button.guest.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
