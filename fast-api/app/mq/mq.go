package main

import (
	"fast-boot/app/mq/internal/listen"
	"flag"
	"github.com/zeromicro/go-zero/core/service"

	"fast-boot/app/mq/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/mq.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}
