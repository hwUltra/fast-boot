package test

import (
	"encoding/json"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/hwUltra/fb-tools/mq"
	"testing"
	"time"
)

func Test_Enqueue(t *testing.T) {
	err := mq.DfSend("192.168.3.88:16379", globalkey.AsyncTaskOrder, map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"})
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}

}

func Test_Enqueue_delay(t *testing.T) {

	payload := map[string]interface{}{"OrderNo": "6666", "OrderMsg": "success ok"}
	err := mq.DfSendAfter("192.168.3.88:16379", globalkey.AsyncTaskOrder, payload, 10*time.Second)
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
}

func Test_EnqueueOther(t *testing.T) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "192.168.3.88:16379"})
	payload := map[string]interface{}{"OrderNo": "6666", "OrderMsg": "success ok"}
	b, _ := json.Marshal(payload)
	task := asynq.NewTask("async_task_order", b)
	// 10秒超时，最多重试3次，20秒后过期
	res, err := client.Enqueue(
		task,
		asynq.MaxRetry(3),
		asynq.Timeout(10*time.Second),
		asynq.Deadline(time.Now().Add(20*time.Second)))
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
	fmt.Printf("Enqueued Result: %+v\n", res)
}
