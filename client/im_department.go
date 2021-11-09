package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImDepartmentGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.department.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDepartmentColleaguesList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.department.colleagues.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDepartmentColleaguesGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.department.colleagues.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDepartmentManagersGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.department.managers.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDepartmentEmployeesGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.department.employees.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
