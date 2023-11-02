package handler

import (
	"context"
	"fast-boot/app/mq/internal/config"
	"fast-boot/app/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// Mqs 返回所有消费者
func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	services = append(services, AsynqMqs(ctx, svcContext)...)

	return services
}

// AsynqMqs  异步队列
func AsynqMqs(ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//监听延迟队列
		NewAsynqJob(ctx, svcContext),
		//监听定时任务
		NewSchedulerJob(ctx, svcContext),
	}

}
