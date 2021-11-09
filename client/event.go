package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) EventBind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.bind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EventUnbind(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.unbind", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EventGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EventOfflineGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.offline.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EventOfflineClear(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.offline.clear", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EventOfflineError(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.offline.error", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EventOfflineList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.offline.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) EventTest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("event.test", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
