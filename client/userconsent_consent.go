package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) UserconsentConsentAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("userconsent.consent.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
