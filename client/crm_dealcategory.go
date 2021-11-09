package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmDealcategoryFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryStatus(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.status", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryStageList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.stage.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryDefaultGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.default.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmDealcategoryDefaultSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.dealcategory.default.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
