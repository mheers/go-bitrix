package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) LandingBlockClonecard(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.clonecard", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockAddcard(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.addcard", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockRemovecard(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.removecard", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockUpdatecards(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.updatecards", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockChangenodename(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.changenodename", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockChangeanchor(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.changeanchor", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockUpdatenodes(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.updatenodes", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockUpdatestyles(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.updatestyles", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockUpdateattrs(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.updateattrs", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockGetcontent(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.getcontent", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockUpdatecontent(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.updatecontent", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockGetlist(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.getlist", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockGetbyid(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.getbyid", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockGetmanifestfile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.getmanifestfile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockGetcontentfromrepository(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.getcontentfromrepository", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockGetrepository(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.getrepository", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingBlockUploadfile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.block.uploadfile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
