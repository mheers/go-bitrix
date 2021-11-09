package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TelephonyCallAttachtranscription(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.call.attachtranscription", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
