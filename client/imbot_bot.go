package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImbotBotList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.bot.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
