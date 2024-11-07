package config

import (
	"github.com/hwUltra/fb-tools/uploadx"
	"github.com/zeromicro/go-zero/core/stores/cache"
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
	CacheConf cache.CacheConf
	RedisConf redis.RedisConf
	OSSConf   uploadx.OSSConf
	WebSocket struct {
		Enable bool
		Path   string
	}
}
