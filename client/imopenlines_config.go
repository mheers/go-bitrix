package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImopenlinesConfigPathGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.config.path.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesConfigGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.config.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesConfigListGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.config.list.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesConfigUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.config.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesConfigAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.config.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesConfigDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.config.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
