package handler

import (
	"context"
	"fast-boot/app/mq/internal/job"
	"fast-boot/app/mq/internal/svc"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type AsyncJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsyncJob(ctx context.Context, svcCtx *svc.ServiceContext) *AsyncJob {
	return &AsyncJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsyncJob) Start() {

	fmt.Println("AsyncJob start ")

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
	mux.HandleFunc(job.AsyncTaskOrder, job.HandleTaskOrder)
	mux.HandleFunc(job.AsyncTaskGift, job.HandleTaskGift)

	mux.HandleFunc(job.SchedulerTestTask, job.HandleSchedulerTest)
	mux.HandleFunc(job.SchedulerMailTask, job.HandleSchedulerMail)

	if err := srv.Run(mux); err != nil {
		logx.Error("could not run server: %v", err)
	}
}

func (l *AsyncJob) Stop() {
	fmt.Println("AsyncJob stop")
}
