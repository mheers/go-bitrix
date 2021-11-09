package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) Sonet_group_subjectGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group_subject.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_group_subjectAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group_subject.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_group_subjectUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group_subject.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Sonet_group_subjectDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("sonet_group_subject.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
