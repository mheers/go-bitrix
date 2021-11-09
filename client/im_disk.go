package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImDiskFolderGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.disk.folder.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDiskFileCommit(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.disk.file.commit", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDiskFileDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.disk.file.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImDiskFileSave(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.disk.file.save", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
