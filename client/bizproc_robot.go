package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) BizprocRobotAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.robot.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocRobotUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.robot.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocRobotDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.robot.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) BizprocRobotList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("bizproc.robot.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
