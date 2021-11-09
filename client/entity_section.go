package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) EntitySectionAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.section.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntitySectionGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.section.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntitySectionUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.section.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EntitySectionDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("entity.section.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
