package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CalendarUserSettingsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.user.settings.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarUserSettingsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.user.settings.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
