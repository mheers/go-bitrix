package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) PlacementList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("placement.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) PlacementBind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("placement.bind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) PlacementUnbind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("placement.unbind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) PlacementGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("placement.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
