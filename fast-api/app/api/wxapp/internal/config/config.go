package config

import (
	"github.com/hwUltra/fb-tools/pay"
	"github.com/hwUltra/fb-tools/sms"
	"github.com/hwUltra/fb-tools/wechat"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	WxMini    wechat.WxMiniConf
	WxPay     pay.WxConf
	Sms       sms.VCodeConf
	RedisConf redis.RedisConf
	CacheConf cache.CacheConf
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

func PullConfig() Config {
	const (
		EtcdHost = "192.168.3.88:12379"
		EtcdKey  = "fa-api-wxapp"
	)
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{EtcdHost},
		Key:   EtcdKey,
	})

	cc := configurator.MustNewConfigCenter[Config](configurator.Config{
		Type: "yaml",
	}, ss)

	// 获取配置
	v, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}
	// 监听配置
	cc.AddListener(func() {
		v, err := cc.GetConfig()
		if err != nil {
			panic(err)
		}
		//这个地方要写 触发配置变化后 需要处理的操作
		println("config changed:", v.Name)
	})

	return v
}
