package handler

import (
	"context"
	"fast-boot/app/mq/internal/logic"
	"fast-boot/app/mq/internal/svc"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type AsynqJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqJob(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqJob {
	return &AsynqJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsynqJob) Start() {

	fmt.Println("AsynqJob start ")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()

	mux.HandleFunc("testTask", logic.TestJob)
	mux.HandleFunc("orderTask", logic.OrderJob)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func (l *AsynqJob) Stop() {
	fmt.Println("AsynqJob stop")
}
