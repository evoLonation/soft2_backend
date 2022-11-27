package logic

import (
	errorx "Ingoland/backend/mall/common"
	"Ingoland/backend/mall/service/user/model"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"Ingoland/backend/mall/service/user/api/internal/svc"
	"Ingoland/backend/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *RegisterLogic) getJwtToken(secretKey string, iat, seconds, UserId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["UserId"] = UserId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserModel.FindOneByLoginId(l.ctx, req.LoginId)
	if err == nil {
		return nil, errorx.NewCodeError(1, "用户名已存在")
	}
	newUser := model.User{
		LoginId:  req.LoginId,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	newUser.UserId, err = res.LastInsertId()
	now := time.Now().Unix()
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, newUser.UserId)
	return &types.RegisterResponse{
		Code:     0,
		UserId:   newUser.UserId,
		Token:    jwtToken,
		NickName: req.Nickname,
	}, nil
}
