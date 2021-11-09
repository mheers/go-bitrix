package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImopenlinesSessionIntercept(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.session.intercept", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
