package config

import (
	"fast-boot/common/miniox"
	"fast-boot/common/pay"
	"fast-boot/common/sms"
	"fast-boot/common/wechat"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	WxMini wechat.WxMiniConf
	WxPay  pay.WxConf
	Sms    sms.VCodeConf
	Redis  redis.RedisConf
	MinioX miniox.Conf
	Auth   struct {
		AccessSecret string
		AccessExpire int64
	}
	UmsRpc zrpc.RpcClientConf
}
