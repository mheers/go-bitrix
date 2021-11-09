package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmTimelineCommentFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.comment.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineCommentAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.comment.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineCommentList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.comment.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineCommentGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.comment.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineCommentDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.comment.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineCommentUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.comment.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineBindingsFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.bindings.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineBindingsList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.bindings.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineBindingsBind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.bindings.bind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmTimelineBindingsUnbind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.timeline.bindings.unbind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
