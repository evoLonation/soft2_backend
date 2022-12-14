package main

import (
	"flag"
	"fmt"
	"soft2_backend/common"
	"soft2_backend/service/file/filecommon"

	"soft2_backend/service/file/api/internal/config"
	"soft2_backend/service/file/api/internal/handler"
	"soft2_backend/service/file/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/file-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 自定义错误
	common.InitHttpErrorHandler()

	filecommon.InitFile()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
