package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"soft2_backend/service/file/rpc/types/file"
	"soft2_backend/service/user/model"
	"time"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

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
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}
	if userInfo.Password != req.PassWord {
		return nil, errors.New("密码错误")
	}
	now := time.Now().Unix()
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.UserId)
	if err != nil {
		return nil, err
	}
	avatarUrl, _ := l.svcCtx.FileRpc.GetUserAvatar(l.ctx, &file.UserIdReq{Id: userInfo.UserId})
	return &types.LoginResponse{
		Code:      0,
		UserId:    userInfo.UserId,
		Token:     jwtToken,
		NickName:  userInfo.Nickname,
		AvatarUrl: avatarUrl.Url,
	}, nil
}
