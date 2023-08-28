package tasks

import "github.com/nightwriter/go-bitrix/types"

type TasksResponse struct {
	types.Response
	Result AllTasks `json:"result"`
}
