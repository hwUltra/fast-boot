package config

import (
	"fast-boot/common/gormV2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	Gorm      gormV2.Conf
	JWT       struct {
		AccessSecret string
		AccessExpire int64
	}
	Salt string
}
