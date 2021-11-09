package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) ImChatAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatSetowner(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.setowner", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatSetmanager(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.setmanager", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatUpdatecolor(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.updatecolor", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatUpdatetitle(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.updatetitle", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatUpdateavatar(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.updateavatar", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatLeave(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.leave", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatUserAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.user.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatUserDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.user.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatUserList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.user.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatSendtyping(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.sendtyping", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatMute(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.mute", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImChatParentJoin(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("im.chat.parent.join", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
