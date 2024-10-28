package svc

import (
	"fast-boot/app/mq/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//SysRpc sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//SysRpc: sys.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
