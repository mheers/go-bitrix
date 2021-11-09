package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) FaceClientIdentify(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("face.client.identify", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) FaceClientAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("face.client.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) FaceUserIdentify(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("face.user.identify", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) FaceUserAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("face.user.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) FaceUserDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("face.user.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
