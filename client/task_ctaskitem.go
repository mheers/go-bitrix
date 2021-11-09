package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskCtaskitemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetexecutiveuserid(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getexecutiveuserid", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetdata(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getdata", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetdescription(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getdescription", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetfiles(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getfiles", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemAddfile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.addfile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemDeletefile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.deletefile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGettags(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.gettags", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetdependson(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getdependson", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetallowedtaskactions(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getallowedtaskactions", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetallowedactions(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getallowedactions", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemGetallowedtaskactionsasstrings(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.getallowedtaskactionsasstrings", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemDelegate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.delegate", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemStartexecution(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.startexecution", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemPauseexecution(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.pauseexecution", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemDefer(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.defer", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemComplete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.complete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemRenew(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.renew", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemApprove(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.approve", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemDisapprove(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.disapprove", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemAddtofavorite(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.addtofavorite", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtaskitemDeletefromfavorite(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctaskitem.deletefromfavorite", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
