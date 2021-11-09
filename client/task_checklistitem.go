package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskChecklistitemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemComplete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.complete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemRenew(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.renew", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemMoveafteritem(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.moveafteritem", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskChecklistitemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.checklistitem.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
