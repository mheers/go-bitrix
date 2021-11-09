package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MobileSettingsTabsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.settings.tabs.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileSettingsEnergySet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.settings.energy.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileSettingsEnergyGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.settings.energy.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileSettingsOtherSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.settings.other.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) MobileSettingsOtherGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.settings.other.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
