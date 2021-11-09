package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ListsAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
