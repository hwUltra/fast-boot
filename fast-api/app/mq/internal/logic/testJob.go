package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"time"
)

type TestTaskPayload struct {
	Sn string
}

func TestJob(ctx context.Context, t *asynq.Task) error {

	fmt.Println("------TestJob start------")

	var p TestTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(errors.New("解析 asynq task payload err"), "testJob payload err:%v, payLoad:%+v", err, t.Payload())
	}

	fmt.Println(fmt.Printf("%s -> Sn:%s}", time.Now().String(), p.Sn))

	fmt.Println("------TestJob end------")

	return nil

}
