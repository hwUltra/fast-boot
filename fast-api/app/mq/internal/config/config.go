package config

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf
	Redis redis.RedisConf
	//SendWxMiniTplMessageConf kq.KqConf
	//SysRpc zrpc.RpcClientConf
}

func PullConfig() Config {
	const (
		EtcdHost = "192.168.3.88:12379"
		EtcdKey  = "fa-mq"
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
