package handler

import (
	"context"
	"encoding/json"
	"fast-boot/app/mq/internal/logic"
	"fast-boot/app/mq/internal/svc"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

type SchedulerJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSchedulerJob(ctx context.Context, svcCtx *svc.ServiceContext) *SchedulerJob {
	return &SchedulerJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SchedulerJob) Start() {

	fmt.Println("SchedulerJob start ")
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	srv := asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)

	payload, err := json.Marshal(logic.TestTaskPayload{Sn: "msg for sn"})
	if err != nil {
		log.Fatal(err)
	}
	entryID, err := srv.Register("*/2 * * * *", asynq.NewTask("testTask", payload))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func (l *SchedulerJob) Stop() {
	fmt.Println("SchedulerJob stop")
}
