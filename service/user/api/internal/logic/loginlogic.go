package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	errorx "soft2_backend/common"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
	"soft2_backend/service/user/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, UserId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["UserId"] = UserId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserModel.FindOneByLoginId(l.ctx, req.LoginId)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewCodeError(1, "用户名不存在")
	default:
		return nil, err
	}
	if userInfo.Password != req.PassWord {
		return nil, errorx.NewCodeError(2, "密码错误")
	}
	now := time.Now().Unix()
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Printf("~~~~~~\n%s\n~~~~~~~~~", l.ctx.Value("exp"))
	return &types.LoginResponse{
		Code:     0,
		Token:    jwtToken,
		UserId:   userInfo.UserId,
		NickName: userInfo.Nickname,
	}, nil
}
