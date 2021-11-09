package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) BizprocActivityAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.activity.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocActivityUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.activity.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocActivityDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.activity.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocActivityLog(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.activity.log", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocActivityList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.activity.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
