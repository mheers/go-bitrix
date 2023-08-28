package departments

import "github.com/nightwriter/go-bitrix/types"

type DepartmentResponse struct {
	types.Response
	Result []Department `json:"result"`
}
