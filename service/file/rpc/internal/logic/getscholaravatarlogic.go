package logic

import (
	"context"
	"errors"
	"soft2_backend/service/file/filecommon"

	"soft2_backend/service/file/rpc/internal/svc"
	"soft2_backend/service/file/rpc/types/file"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScholarAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScholarAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScholarAvatarLogic {
	return &GetScholarAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScholarAvatarLogic) GetScholarAvatar(in *file.ScholarIdReq) (*file.UrlReply, error) {
	//todo 通过学者id得到userid
	var userId int64
	return nil, errors.New("根据学者id获取用户id的接口还没有实装")
	//res, err := l.svcCtx.Apply.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: in.Id})
	//userId = res.UserId
	userAvatar, err := l.svcCtx.UserAvatarModel.FindOne(l.ctx, userId)
	err = filecommon.SqlErrorCheck(err)
	if err != nil && err != filecommon.NoRowError {
		return nil, err
	}
	if err == filecommon.NoRowError {
		return &file.UrlReply{Url: filecommon.GetDefaultAvatarUrl()}, nil
	} else {
		return &file.UrlReply{Url: filecommon.GetUrl(userAvatar.FileName)}, nil
	}
}
