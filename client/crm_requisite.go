package client

import "github.com/nightwriter/go-bitrix/types"

func (c *Client) CrmRequisiteFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteUserfieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.userfield.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteUserfieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.userfield.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteUserfieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.userfield.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteUserfieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.userfield.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteUserfieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.userfield.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetCountries(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.countries", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFieldFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.field.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFieldAvailabletoadd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.field.availabletoadd", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFieldAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.field.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFieldGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.field.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFieldList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.field.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFieldUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.field.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisitePresetFieldDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.preset.field.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteBankdetailFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.bankdetail.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteBankdetailAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.bankdetail.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteBankdetailGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.bankdetail.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteBankdetailList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.bankdetail.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteBankdetailUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.bankdetail.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteBankdetailDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.bankdetail.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteLinkFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.link.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteLinkList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.link.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteLinkGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.link.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteLinkRegister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.link.register", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmRequisiteLinkUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.requisite.link.unregister", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
