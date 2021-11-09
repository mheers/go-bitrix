package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) Download(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("download", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Upload(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("upload", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
