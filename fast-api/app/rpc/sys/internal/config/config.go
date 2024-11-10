package config

import (
	"github.com/hwUltra/fb-tools/gormx"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
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

func PullConfig() Config {
	const (
		EtcdHost = "192.168.3.88:12379"
		EtcdKey  = "fa-rpc-sys"
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
