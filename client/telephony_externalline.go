package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TelephonyExternallineAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalline.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternallineUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalline.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternallineDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalline.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternallineGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalline.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
