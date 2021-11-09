package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) VoximplantCallbackStart(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.callback.start", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantInfocallStartwithtext(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.infocall.startwithtext", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantInfocallStartwithsound(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.infocall.startwithsound", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
