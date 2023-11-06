package main

import (
	"fast-boot/app/ws/internal/config"
	"fast-boot/app/ws/internal/service"
	"fast-boot/common/websocket/core"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

//var (
//	port    = flag.Int("port", 3333, "the port to listen")
//	timeout = flag.Int64("timeout", 0, "timeout of milliseconds")
//	cpu     = flag.Int64("cpu", 500, "cpu threshold")
//)

var configFile = flag.String("f", "./etc/ws.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	hub := core.CreateHubFactory()
	go hub.Run()

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			if r.Method != "GET" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			http.ServeFile(w, r, "home.html")
		},
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if ws, ok := service.OnOpen(hub, w, r, c); ok {
				ws.OnMessage(r.Context())
			}
		},
	})

	server.Start()
}
