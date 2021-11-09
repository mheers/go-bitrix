package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) VoximplantUrlGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.url.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantSipGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.sip.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantSipAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.sip.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantSipUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.sip.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantSipDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.sip.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantSipStatus(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.sip.status", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantSipConnectorStatus(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.sip.connector.status", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantStatisticGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.statistic.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantLineOutgoingSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.line.outgoing.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantLineOutgoingGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.line.outgoing.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantLineOutgoingSipSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.line.outgoing.sip.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantLineGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.line.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantTtsVoicesGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.tts.voices.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantUserGetdefaultlineid(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.user.getdefaultlineid", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantUserActivatephone(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.user.activatephone", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantUserDeactivatephone(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.user.deactivatephone", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantAuthorizationGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.authorization.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantAuthorizationSignonetimekey(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.authorization.signonetimekey", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantAuthorizationOnerror(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.authorization.onerror", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallInit(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.init", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallStartwithdevice(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.startwithdevice", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallHangupdevice(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.hangupdevice", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallSendwait(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.sendwait", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallSendready(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.sendready", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallAnswer(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.answer", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallSkip(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.skip", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallHold(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.hold", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallUnhold(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.unhold", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallStartviarest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.startviarest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallIntercept(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.intercept", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) VoximplantCallSavecomment(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("voximplant.call.savecomment", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
