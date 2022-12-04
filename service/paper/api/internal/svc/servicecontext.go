package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/file/rpc/fileclient"
	"soft2_backend/service/paper/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	FileRpc fileclient.File
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		FileRpc: fileclient.NewFile(zrpc.MustNewClient(c.FileRpc)),
	}
}
