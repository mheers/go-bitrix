package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingLandingGetpreview(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.getpreview", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingGetpublicurl(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.getpublicurl", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingGetadditionalfields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.getadditionalfields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingPublication(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.publication", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingUnpublic(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.unpublic", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingAddblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.addblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingDeleteblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.deleteblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingMarkdeletedblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.markdeletedblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingMarkundeletedblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.markundeletedblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingUpblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.upblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingDownblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.downblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingShowblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.showblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingHideblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.hideblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingCopyblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.copyblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingMoveblock(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.moveblock", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingRemoveentities(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.removeentities", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingAddbytemplate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.addbytemplate", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingCopy(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.copy", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingMarkdelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.markdelete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingLandingMarkundelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.landing.markundelete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
