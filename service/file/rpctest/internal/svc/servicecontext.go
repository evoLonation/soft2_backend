package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/file/rpc/fileclient"
	"soft2_backend/service/file/rpc/types/file"
	"soft2_backend/service/file/rpctest/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc file.FileClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FileRpc: fileclient.NewFile(zrpc.MustNewClient(c.FileRpc)),
	}
}
