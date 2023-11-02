package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"time"
)

type OrderTaskPayload struct {
	OrderNo  string
	OrderMsg string
}

func OrderJob(ctx context.Context, t *asynq.Task) error {

	fmt.Println("------OrderJob start------")

	var p OrderTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(errors.New("解析 asynq task payload err"), "payload err:%v, payLoad:%+v", err, t.Payload())
	}

	fmt.Println(fmt.Printf("%s -> OrderNo:%s,OrderMsg:%s}", time.Now().String(), p.OrderNo, p.OrderMsg))

	fmt.Println("------OrderJob end------")
	return nil

}
