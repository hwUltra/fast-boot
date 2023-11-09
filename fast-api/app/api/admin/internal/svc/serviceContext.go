package svc

import (
	"fast-boot/app/api/admin/internal/config"
	"fast-boot/app/api/admin/internal/middleware"
	"fast-boot/app/rpc/sys/sys"
	"fast-boot/app/rpc/ums/ums"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	SysRpc      sys.Sys
	UmsRpc      ums.Ums
	CheckUrl    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.MustNewRedis(c.Redis)
	return &ServiceContext{
		Config:      c,
		RedisClient: newRedis,
		SysRpc:      sys.NewSys(zrpc.MustNewClient(c.SysRpc)),
		UmsRpc:      ums.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		CheckUrl:    middleware.NewCheckUrlMiddleware(newRedis).Handle,
	}
}
