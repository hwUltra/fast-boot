package svc

import (
	"fast-boot/app/api/admin/internal/config"
	"fast-boot/app/rpc/pms/pms"
	"fast-boot/app/rpc/sys/sys"
	"fast-boot/app/rpc/ums/ums"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	SysRpc sys.Sys
	UmsRpc ums.Ums
	PmsRpc pms.Pms
	//CheckUrl rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		SysRpc: sys.NewSys(zrpc.MustNewClient(c.SysRpc)),
		UmsRpc: ums.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		PmsRpc: pms.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		//CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
	}
}
