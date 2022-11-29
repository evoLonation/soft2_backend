package main

import (
	"flag"
	"fmt"
	"soft2_backend/common"

	"soft2_backend/service/file/rpctest/internal/config"
	"soft2_backend/service/file/rpctest/internal/handler"
	"soft2_backend/service/file/rpctest/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/template.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	common.InitHttpErrorHandler()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
