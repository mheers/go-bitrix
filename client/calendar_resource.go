package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CalendarResourceList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.resource.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarResourceAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.resource.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarResourceUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.resource.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarResourceDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.resource.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarResourceBookingList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.resource.booking.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
