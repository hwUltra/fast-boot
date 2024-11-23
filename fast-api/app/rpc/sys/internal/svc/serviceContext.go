package svc

import (
	"fast-boot/app/rpc/sys/internal/config"
	"github.com/hwUltra/fb-tools/gormx"
)

type ServiceContext struct {
	Config     config.Config
	GormClient *gormx.Client
	//Rsc        *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:     c,
		GormClient: gormx.MustNewClient(c.Gorm),
		//Rsc:        redis.MustNewRedis(c.RedisConf),
	}
}
