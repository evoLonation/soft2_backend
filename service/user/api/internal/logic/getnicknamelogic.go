package logic

import (
	"context"
	"errors"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNickNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNickNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNickNameLogic {
	return &GetNickNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNickNameLogic) GetNickName(req *types.GetNickNameRequest) (resp *types.GetNickNameResponse, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err == model.ErrNotFound {
		return nil, errors.New("用户不存在")
	}
	return &types.GetNickNameResponse{NickName: user.Nickname, Email: user.Email}, nil
}
