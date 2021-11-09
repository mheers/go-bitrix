package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) DiskFileGetfields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.getfields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileCopyto(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.copyto", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileMoveto(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.moveto", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileRename(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.rename", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileMarkdeleted(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.markdeleted", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileRestore(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.restore", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileUploadversion(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.uploadversion", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileGetexternallink(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.getexternallink", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileGetversions(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.getversions", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileRestorefromversion(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.restorefromversion", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFileListallowedoperations(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.file.listallowedoperations", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
