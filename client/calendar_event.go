package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CalendarEventGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.event.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarEventAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.event.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarEventUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.event.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarEventDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.event.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarEventGetNearest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.event.get.nearest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarEventGetbyid(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.event.getbyid", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
