package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ListsFieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.field.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsFieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.field.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsFieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.field.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsFieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.field.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsFieldTypeGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.field.type.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
