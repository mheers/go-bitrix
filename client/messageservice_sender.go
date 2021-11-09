package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MessageserviceSenderAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("messageservice.sender.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MessageserviceSenderDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("messageservice.sender.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MessageserviceSenderList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("messageservice.sender.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
