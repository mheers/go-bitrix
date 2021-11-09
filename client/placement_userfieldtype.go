package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) UserfieldtypeList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("userfieldtype.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserfieldtypeAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("userfieldtype.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserfieldtypeUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("userfieldtype.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) UserfieldtypeDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("userfieldtype.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
