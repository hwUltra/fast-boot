package main

import (
	"fast-boot/app/api/wxapp/internal/config"
	"fast-boot/app/api/wxapp/internal/handler"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/wsservice"
	"flag"
	"fmt"
	"github.com/hwUltra/fb-tools/result"
	"github.com/hwUltra/fb-tools/wsCore"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

var configFile = flag.String("f", "etc/wxapp.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	//server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		// Unauthorized
		httpx.WriteJson(w, http.StatusOK, result.Error(401, "认证失败"))
	}))

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	//添加websocket支持
	if c.WebSocket.Enable == true {
		wsHub := wsCore.CreateHubFactory()
		go wsHub.Run()
		server.AddRoute(rest.Route{
			Method: http.MethodGet,
			Path:   c.WebSocket.Path,
			Handler: func(w http.ResponseWriter, r *http.Request) {
				if ws, ok := wsservice.OnOpen(wsHub, w, r, c); ok {
					ws.OnMessage(r.Context())
				}
			},
		})
		fmt.Printf("Starting websocket server at %s:%d%s...\n", c.Host, c.Port, c.WebSocket.Path)
	}

	server.Start()
}
