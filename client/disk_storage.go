package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) DiskStorageGetfields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.getfields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageRename(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.rename", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageGettypes(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.gettypes", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageAddfolder(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.addfolder", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageGetchildren(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.getchildren", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageUploadfile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.uploadfile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskStorageGetforapp(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.storage.getforapp", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
