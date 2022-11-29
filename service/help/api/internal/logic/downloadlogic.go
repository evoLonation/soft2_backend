package logic

import (
	"context"
	"soft2_backend/service/file/rpc/types/file"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadReq) (resp *types.DownloadReply, err error) {
	// todo: add your logic here and delete this line
	helpFile, err := l.svcCtx.FileRpc.GetHelpFile(l.ctx, &file.ApplyIdReq{
		Id: req.RequestId,
	})
	if err != nil {
		return nil, err
	}
	return &types.DownloadReply{
		Url: helpFile.Url,
	}, nil
}
