package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TimemanTimecontrolSettingsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.timecontrol.settings.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanTimecontrolSettingsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.timecontrol.settings.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanTimecontrolReportAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.timecontrol.report.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanTimecontrolReportsSettingsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.timecontrol.reports.settings.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanTimecontrolReportsUsersGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.timecontrol.reports.users.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanTimecontrolReportsGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.timecontrol.reports.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TimemanTimecontrolReport(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("timeman.timecontrol.report", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
