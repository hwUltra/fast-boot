package svc

import (
	"fast-boot/app/api/wxapp/internal/config"
	"fast-boot/app/rpc/ums/ums"
	"github.com/hwUltra/fb-tools/wechat"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	WxMiniApp   *wechat.MiniTool
	UmsRpc      ums.Ums
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.RedisConf),
		WxMiniApp:   wechat.NewMini(c.WxMini),
		UmsRpc:      ums.NewUms(zrpc.MustNewClient(c.UmsRpc)),
	}
}
