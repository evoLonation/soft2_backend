package main

import (
	"flag"
	"fmt"
	"soft2_backend/common"

	"soft2_backend/service/user/api/internal/config"
	"soft2_backend/service/user/api/internal/handler"
	"soft2_backend/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//这里是自定义错误begin
	common.InitHttpErrorHandler()

	//这里是自定义错误end
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
