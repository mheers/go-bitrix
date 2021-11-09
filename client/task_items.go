package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) TaskItemsGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.items.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
