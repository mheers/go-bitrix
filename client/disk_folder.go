package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) DiskFolderGetfields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.getfields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderGetchildren(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.getchildren", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderAddsubfolder(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.addsubfolder", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderCopyto(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.copyto", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderMoveto(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.moveto", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderRename(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.rename", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderDeletetree(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.deletetree", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderMarkdeleted(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.markdeleted", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderRestore(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.restore", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderUploadfile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.uploadfile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderGetexternallink(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.getexternallink", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderSharetouser(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.sharetouser", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) DiskFolderListallowedoperations(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("disk.folder.listallowedoperations", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
