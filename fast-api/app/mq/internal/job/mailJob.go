package job

import (
	"context"
	"encoding/json"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

const SchedulerMailTask = globalkey.SchedulerMailTask

type SchedulerMailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// NewSchedulerMail 创建异步任务的函数
func NewSchedulerMail(sPayload SchedulerMailPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(sPayload)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(SchedulerMailTask, payload)
	return task, err
}

// HandleSchedulerMail 处理异步任务的函数
func HandleSchedulerMail(ctx context.Context, task *asynq.Task) error {
	payload := SchedulerMailPayload{}
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	// TODO: 模拟处理
	fmt.Println("SchedulerMail begin!")
	fmt.Printf("%s - Sending to %s, subject: %s, body: %s\n", time.Now().String(), payload.To, payload.Subject, payload.Body)
	fmt.Println("SchedulerMail end")
	return nil
}
