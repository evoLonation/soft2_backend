package logic

import (
	"context"
	"soft2_backend/service/file/rpc/fileclient"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarGetAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarGetAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarGetAvatarLogic {
	return &ScholarGetAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarGetAvatarLogic) ScholarGetAvatar(req *types.ScholarGetAvatarRequest) (resp *types.ScholarGetAvatarResponse, err error) {
	// todo: add your logic here and delete this line
	avatarUrl, err := l.svcCtx.FileRpc.GetScholarAvatar(l.ctx, &fileclient.ScholarIdReq{
		Id: req.ScholarId,
	})
	if err != nil {
		resp = &types.ScholarGetAvatarResponse{
			Code: -1,
			Url: "",
		}
	} else {
		resp = &types.ScholarGetAvatarResponse{
			Code: 0,
			Url: avatarUrl.Url,
		}
	}
	return resp, nil
}
