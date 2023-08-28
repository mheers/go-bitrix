package activity_stream

import "github.com/nightwriter/go-bitrix/types"

type ActivityStreamResponse struct {
	types.Response
	Result []ActivityStream `json:"result"`
}
