package job

import (
	"context"
	"encoding/json"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

const SchedulerTestTask = globalkey.SchedulerTestTask

type SchedulerTestPayload struct {
	Sn string
}

// NewSchedulerTest 创建异步任务的函数
func NewSchedulerTest(sPayload SchedulerTestPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(sPayload)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(SchedulerTestTask, payload)
	return task, err
}

// HandleSchedulerTest 处理异步任务的函数
func HandleSchedulerTest(ctx context.Context, task *asynq.Task) error {
	payload := SchedulerTestPayload{}
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	// TODO: 模拟处理
	fmt.Println("SchedulerTest begin!")
	fmt.Printf("%s - sn: %s\n", time.Now().String(), payload.Sn)
	fmt.Println("SchedulerTest end")

	return nil
}
