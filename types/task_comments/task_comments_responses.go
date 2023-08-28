package task_comments

import "github.com/nightwriter/go-bitrix/types"

type TaskCommentsResponse struct {
	Result             []TaskComments `json:"result"`
	types.ResponseTime `json:"time"`
}
