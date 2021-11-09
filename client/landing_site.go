package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingSiteGetadditionalfields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.getadditionalfields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteGetpublicurl(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.getpublicurl", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteMarkdelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.markdelete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteMarkundelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.markundelete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSitePublication(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.publication", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteUnpublic(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.unpublic", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteFullexport(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.fullexport", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteSetrights(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.setrights", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteGetrights(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.getrights", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingSiteSetscope(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.site.setscope", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
