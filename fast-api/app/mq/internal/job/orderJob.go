package job

import (
	"context"
	"encoding/json"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hibiken/asynq"
)

const AsyncTaskOrder = globalkey.AsyncTaskOrder

type TaskOrderPayload struct {
	OrderNo  string
	OrderMsg string
}

// NewTaskOrder 创建异步任务的函数
func NewTaskOrder(asyncTask TaskOrderPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(asyncTask)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(AsyncTaskOrder, payload)
	return task, err
}

// HandleTaskOrder 处理异步任务的函数
func HandleTaskOrder(ctx context.Context, task *asynq.Task) error {
	payload := TaskOrderPayload{}
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	// TODO: 模拟处理
	fmt.Println("OrderTask begin!")
	fmt.Printf("OrderNo to %s,  OrderMsg: %s\n", payload.OrderNo, payload.OrderMsg)
	fmt.Println("OrderTask end!")
	return nil
}
