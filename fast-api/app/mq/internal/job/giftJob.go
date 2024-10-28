package job

import (
	"context"
	"encoding/json"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hibiken/asynq"
)

const AsyncTaskGift = globalkey.AsyncTaskGift

type TaskGiftPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// NewTaskGift 创建异步任务的函数
func NewTaskGift(tPayload TaskGiftPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(tPayload)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(AsyncTaskGift, payload)
	return task, err
}

// HandleTaskGift 处理异步任务的函数
func HandleTaskGift(ctx context.Context, task *asynq.Task) error {
	payload := TaskGiftPayload{}
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	// TODO: 模拟处理
	fmt.Println("HandleAsyncGiftTask begin!")
	fmt.Printf("Sending to %s, subject: %s, body: %s\n", payload.To, payload.Subject, payload.Body)
	fmt.Println("HandleAsyncGiftTask end!")
	return nil
}
