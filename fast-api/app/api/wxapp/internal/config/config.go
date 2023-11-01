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
	WxMiniConf wechat.WxMiniConf
	WxPayConf  pay.WxConf
	SmsConf    sms.Conf
	RedisConf  redis.RedisConf
	MinioX     miniox.Conf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	UmsRpc zrpc.RpcClientConf
	//CartRpc    zrpc.RpcClientConf
	//ProductRpc zrpc.RpcClientConf
	//OrderRpc   zrpc.RpcClientConf
	//PayRpc     zrpc.RpcClientConf
}
