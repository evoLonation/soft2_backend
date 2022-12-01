package main

import (
	"flag"
	"fmt"
	"soft2_backend/service/paper/api/internal/config"
	"soft2_backend/service/paper/api/internal/handler"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/database"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/paper-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	database.Init()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
