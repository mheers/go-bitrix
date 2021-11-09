package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImbotChatAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatSetowner(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.setowner", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatSetmanager(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.setmanager", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatUpdatecolor(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.updatecolor", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatUpdatetitle(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.updatetitle", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatUpdateavatar(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.updateavatar", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatLeave(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.leave", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatUserAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.user.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatUserDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.user.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatUserList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.user.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImbotChatSendtyping(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imbot.chat.sendtyping", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
