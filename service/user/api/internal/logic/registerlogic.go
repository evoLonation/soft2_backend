package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"soft2_backend/service/file/rpc/types/file"
	"soft2_backend/service/help/rpc/helpclient"
	message2 "soft2_backend/service/message/rpc/types/message"
	"soft2_backend/service/user/model"
	"time"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

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
		return nil, errors.New("用户名已注册")
	}
	newUser := model.User{
		LoginId:  req.LoginId,
		Password: req.Password,
		Nickname: req.Nickname,
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	newUser.UserId, err = res.LastInsertId()
	now := time.Now().Unix()
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, newUser.UserId)
	//调用文献互助rpc
	_, _ = l.svcCtx.HelpRpc.RegisterUser(l.ctx, &helpclient.IdReq{
		Id: newUser.UserId,
	})
	avatarUrl, _ := l.svcCtx.FileRpc.GetUserAvatar(l.ctx, &file.UserIdReq{Id: newUser.UserId})
	_, err = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  newUser.UserId,
		Content:     "欢迎使用学术交流平台",
		MessageType: 0,
	})
	fmt.Printf("``````\n%s\n", err)
	return &types.RegisterResponse{
		Code:      0,
		UserId:    newUser.UserId,
		Token:     jwtToken,
		NickName:  req.Nickname,
		AvatarUrl: avatarUrl.Url,
	}, nil
}
