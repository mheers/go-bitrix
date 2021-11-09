package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) BizprocProviderAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.provider.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocProviderDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.provider.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocProviderList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.provider.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
