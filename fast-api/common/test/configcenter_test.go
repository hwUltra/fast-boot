package test

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"testing"
)

func TestNewConfigCenter(t *testing.T) {
	// 创建 configurator
	// 配置结构定义
	//type TestSt struct {
	//	Name string `json:"name"`
	//}
	//// 创建 etcd subscriber
	//ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
	//	Hosts: []string{"192.168.3.88:12379"}, // etcd 地址
	//	Key:   "test1",                        // 配置key
	//})
	//cc := configurator.MustNewConfigCenter[TestSt](configurator.Config{
	//	Type: "json", // 配置值类型：json,yaml,toml
	//}, ss)

}

func Test_ConfigCenter(t *testing.T) {
	// 配置结构定义
	type TestSt struct {
		Name string `json:"name"`
	}
	// 创建 etcd subscriber
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{"192.168.3.88:12379"}, // etcd 地址
		Key:   "fast-boot",                    // 配置key
	})

	// 创建 configurator
	cc := configurator.MustNewConfigCenter[TestSt](configurator.Config{
		Type: "json", // 配置值类型：json,yaml,toml
	}, ss)

	// 获取配置
	// 注意: 配置如果发生变更，调用的结果永远获取到最新的配置
	v, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}
	println(v.Name)

	// 如果想监听配置变化，可以添加 listener
	cc.AddListener(func() {
		v, err := cc.GetConfig()
		if err != nil {
			panic(err)
		}
		println(v.Name)
	})

	select {}
}
