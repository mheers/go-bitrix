package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ListsSectionAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.section.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsSectionGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.section.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsSectionUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.section.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ListsSectionDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.section.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
