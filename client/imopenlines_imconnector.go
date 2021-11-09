package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImconnectorActivate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.activate", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorStatus(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.status", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorConnectorDataSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.connector.data.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorSendMessages(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.send.messages", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorUpdateMessages(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.update.messages", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorDeleteMessages(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.delete.messages", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorSendStatusDelivery(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.send.status.delivery", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorSendStatusReading(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.send.status.reading", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImconnectorSetError(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imconnector.set.error", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
