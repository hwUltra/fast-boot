package config

import (
	"github.com/hwUltra/fb-tools/miniox"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//会员
	SysRpc zrpc.RpcClientConf
	UmsRpc zrpc.RpcClientConf
	PmsRpc zrpc.RpcClientConf
	Auth   struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis     redis.RedisConf
	MinioX    miniox.Conf
	WebSocket struct {
		Path string
	}
}
