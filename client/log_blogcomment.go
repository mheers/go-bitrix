package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LogBlogcommentAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogcomment.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogcommentUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogcomment.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogBlogcommentDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.blogcomment.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
