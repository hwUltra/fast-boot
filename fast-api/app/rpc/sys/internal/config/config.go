package config

import (
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	GormConf  gormx.Conf
	CacheConf cache.CacheConf
	JWT       struct {
		AccessSecret string
		AccessExpire int64
	}
	Salt string
}
