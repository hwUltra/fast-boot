package listen

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

	//kq ：消息队列.
	//services = append(services, KqMqs(c, ctx, svcContext)...)
	//asynq ： 延迟队列、定时任务
	services = append(services, AsynqMqs(c, ctx, svcContext)...)

	return services
}
