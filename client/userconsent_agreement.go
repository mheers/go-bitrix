package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) UserconsentAgreementList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("userconsent.agreement.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserconsentAgreementText(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("userconsent.agreement.text", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
