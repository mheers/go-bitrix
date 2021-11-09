package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TelephonyExternalcallSearchcrmentities(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalcall.searchcrmentities", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternalcallRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalcall.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternalcallFinish(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalcall.finish", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternalcallShow(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalcall.show", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternalcallHide(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalcall.hide", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TelephonyExternalcallAttachrecord(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("telephony.externalcall.attachrecord", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
