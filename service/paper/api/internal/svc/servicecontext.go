package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/apply/rpc/applyclient"
	"soft2_backend/service/file/rpc/fileclient"
	"soft2_backend/service/paper/api/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	FileRpc  fileclient.File
	ApplyRpc applyclient.Apply
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		FileRpc:  fileclient.NewFile(zrpc.MustNewClient(c.FileRpc)),
		ApplyRpc: applyclient.NewApply(zrpc.MustNewClient(c.ApplyRpc)),
	}
}
