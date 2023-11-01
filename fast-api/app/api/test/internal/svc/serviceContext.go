package svc

import (
	"fast-boot/app/api/test/internal/config"
	"fast-boot/app/rpc/sys/sys"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	SysRpc sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SysRpc: sys.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
