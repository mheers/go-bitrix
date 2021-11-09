package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MobileIntranetDepartmentsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.intranet.departments.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileIntranetStresslevelSharedataGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.intranet.stresslevel.sharedata.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
