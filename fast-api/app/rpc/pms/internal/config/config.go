package config

import (
	"github.com/hwUltra/fb-tools/gormV2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Gorm      gormV2.Conf
	RedisConf redis.RedisConf
}
