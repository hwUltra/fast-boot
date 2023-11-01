package svc

import (
	"fast-boot/app/rpc/ums/internal/config"
	"fast-boot/common/gormV2"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	GormConn *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormConn, _ := gormV2.GetSqlDriver(c.Gorm)
	return &ServiceContext{
		Config:   c,
		GormConn: gormConn,
	}
}
