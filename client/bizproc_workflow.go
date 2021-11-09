package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) BizprocWorkflowTerminate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.terminate", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocWorkflowStart(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.start", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocWorkflowInstanceList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.instance.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocWorkflowTemplateList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.template.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocWorkflowTemplateAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.template.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocWorkflowTemplateUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.template.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocWorkflowTemplateDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.template.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocWorkflowInstances(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.workflow.instances", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
