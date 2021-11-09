package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskItemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetexecutiveuserid(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getexecutiveuserid", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetdata(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getdata", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetdescription(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getdescription", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetfiles(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getfiles", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemAddfile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.addfile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemDeletefile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.deletefile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGettags(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.gettags", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetdependson(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getdependson", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetallowedtaskactions(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getallowedtaskactions", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetallowedactions(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getallowedactions", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemGetallowedtaskactionsasstrings(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.getallowedtaskactionsasstrings", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemIsactionallowed(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.isactionallowed", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemDelegate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.delegate", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemStartexecution(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.startexecution", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemPauseexecution(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.pauseexecution", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemDefer(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.defer", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemComplete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.complete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemRenew(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.renew", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemApprove(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.approve", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemDisapprove(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.disapprove", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemAddtofavorite(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.addtofavorite", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemDeletefromfavorite(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.deletefromfavorite", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUserfieldGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.userfield.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUserfieldGetfields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.userfield.getfields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskItemUserfieldGettypes(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.item.userfield.gettypes", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
