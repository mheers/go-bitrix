package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) MobileTasksDeadlinesGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.tasks.deadlines.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
