package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CalendarSectionGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.section.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarSectionAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.section.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarSectionUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.section.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarSectionDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.section.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
