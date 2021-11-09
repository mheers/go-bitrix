package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ListsElementAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.element.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsElementGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.element.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsElementUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.element.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsElementDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.element.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsElementGetFileUrl(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.element.get.file.url", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
