package config

import (
	"github.com/hwUltra/fb-tools/miniox"
	"github.com/hwUltra/fb-tools/pay"
	"github.com/hwUltra/fb-tools/sms"
	"github.com/hwUltra/fb-tools/wechat"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	WxMini    wechat.WxMiniConf
	WxPay     pay.WxConf
	Sms       sms.VCodeConf
	Redis     redis.RedisConf
	MinioX    miniox.Conf
	WebSocket struct {
		Enable bool
		Path   string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UmsRpc zrpc.RpcClientConf
}
