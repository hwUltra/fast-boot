package svc

import (
	"fast-boot/app/rpc/sys/internal/config"
	"fast-boot/common/gormV2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	GormConn    *gorm.DB
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	gormConn, _ := gormV2.GetSqlDriver(c.Gorm)
	return &ServiceContext{
		Config:      c,
		GormConn:    gormConn,
		RedisClient: redis.MustNewRedis(c.RedisConf),
	}
}
