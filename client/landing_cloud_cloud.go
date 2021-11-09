package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) Landing_cloudCloudGetrepository(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing_cloud.cloud.getrepository", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Landing_cloudCloudGetdemositelist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing_cloud.cloud.getdemositelist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Landing_cloudCloudGetdemopagelist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing_cloud.cloud.getdemopagelist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Landing_cloudCloudGeturlpreview(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing_cloud.cloud.geturlpreview", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Landing_cloudCloudGetappitems(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing_cloud.cloud.getappitems", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Landing_cloudCloudGetappitemmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing_cloud.cloud.getappitemmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
