package users

import "github.com/nightwriter/go-bitrix/types"

type UsersResponse struct {
	types.Response
	Result []User `json:"result"`
}
