package test

import (
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hwUltra/fb-tools/mq"
	"testing"
	"time"
)

func Test_Enqueue(t *testing.T) {
	client := mq.NewTaskClient(
		mq.ClientConfig{Addr: "192.168.3.88:16379"},
	)

	dispatch, err := client.Dispatch(
		globalkey.AsyncTaskOrder,
		map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"},
	)
	fmt.Println("Test_Enqueue = ", dispatch, err)

}

func Test_Enqueue_delay(t *testing.T) {
	client := mq.NewTaskClient(
		mq.ClientConfig{Addr: "192.168.3.88:16379"},
	)

	dispatch, err := client.Dispatch(
		globalkey.AsyncTaskOrder,
		map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"},
		client.SetProcessIn(10*time.Second),
	)
	fmt.Println("Test_Enqueue_delay = ", dispatch, err)
}

func Test_EnqueueOther(t *testing.T) {
	client := mq.NewTaskClient(
		mq.ClientConfig{Addr: "192.168.3.88:16379"},
	)

	dispatch, err := client.Dispatch(
		globalkey.AsyncTaskOrder,
		map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"},
		client.SetProcessAt(time.Now().Add(2*time.Hour)),
	)
	fmt.Println("Test_EnqueueOther = ", dispatch, err)
}
