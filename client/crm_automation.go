package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmAutomationTrigger(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.automation.trigger", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAutomationTriggerAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.automation.trigger.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAutomationTriggerList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.automation.trigger.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAutomationTriggerDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.automation.trigger.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmAutomationTriggerExecute(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.automation.trigger.execute", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
