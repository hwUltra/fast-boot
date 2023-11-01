package listen

import (
	"context"
	"fast-boot/app/mq/internal/config"
	"fast-boot/app/mq/internal/mqs/deferMq"
	"fast-boot/app/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// AsyncMqs
func AsynqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//监听消费流水状态变更
		//.....
		deferMq.NewAsynqJob(ctx, svcContext),
		deferMq.NewSchedulerJob(ctx, svcContext),
	}

}
