package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LogBlogpostGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogpost.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogpostUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogpost.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogpostAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogpost.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogpostUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogpost.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogpostShare(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogpost.share", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogpostDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogpost.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogpostGetusersImportant(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogpost.getusers.important", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
