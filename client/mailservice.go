package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MailserviceFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mailservice.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MailserviceList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mailservice.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MailserviceGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mailservice.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MailserviceAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mailservice.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MailserviceUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mailservice.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MailserviceDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mailservice.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
