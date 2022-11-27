package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	fmt.Printf("hereherehereherehere%dherehereherehere\n", userId)
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	} //不过应该不会存在上述情况
	return &types.UserInfoResponse{
		Nickname:   res.Nickname,
		Email:      res.Email,
		Requests:   res.Requests,
		Helps:      res.Help,
		Follows:    res.Follows,
		Complaints: res.Complaints,
		Wealth:     res.Money,
	}, nil
}
