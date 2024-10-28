package handler

import (
	"context"
	"fast-boot/app/mq/internal/job"
	"fast-boot/app/mq/internal/svc"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
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

	//启动定时任务1
	testScheduler, _ := job.NewSchedulerTest(job.SchedulerTestPayload{Sn: "msg for sn"})
	entryID, err := srv.Register("*/1 * * * *", testScheduler)
	fmt.Printf("registered an entry: %q\n", entryID)
	//启动定时任务2
	mailScheduler, _ := job.NewSchedulerMail(job.SchedulerMailPayload{To: "msg for sn",
		Subject: "xxx", Body: "bbbb"})
	entryID2, err := srv.Register("@every 30s", mailScheduler)
	fmt.Printf("registered an entry2: %q\n", entryID2)

	//开启
	if err := srv.Run(); err != nil {
		logx.Error("srv run: %v", err)
	}
}

func (l *SchedulerJob) Stop() {
	fmt.Println("SchedulerJob stop")
}
