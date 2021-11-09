package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImSearchUserList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.user.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchUser(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.user", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchChatList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.chat.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchChat(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.chat", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchDepartmentList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.department.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchDepartment(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.department", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchLastGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.last.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchLastAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.last.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImSearchLastDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.search.last.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
