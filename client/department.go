package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) DepartmentFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("department.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DepartmentGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("department.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DepartmentAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("department.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DepartmentUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("department.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DepartmentDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("department.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
