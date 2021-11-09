package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) PullConfigGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("pull.config.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) PullChannelPublicGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("pull.channel.public.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) PullChannelPublicList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("pull.channel.public.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
