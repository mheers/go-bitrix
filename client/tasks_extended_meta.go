package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) Tasks_extendedMetaSetanystatus(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("tasks_extended.meta.setanystatus", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Tasks_extendedMetaOccurinlogsas(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("tasks_extended.meta.occurinlogsas", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
