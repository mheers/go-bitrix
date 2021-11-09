package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImDialogGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.dialog.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDialogMessagesGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.dialog.messages.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDialogUsersGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.dialog.users.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDialogRead(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.dialog.read", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDialogReadall(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.dialog.readall", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDialogUnread(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.dialog.unread", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDialogWriting(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.dialog.writing", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
