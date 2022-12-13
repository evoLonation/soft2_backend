package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"soft2_backend/service/apply/api/internal/svc"
	"soft2_backend/service/apply/api/internal/types"
	"soft2_backend/service/apply/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailIdentifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailIdentifyLogic {
	return &EmailIdentifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailIdentifyLogic) EmailIdentify(req *types.EmailIdentifyRequest) (resp *types.EmailIdentifyResponse, err error) {
	verifyCode, err := l.svcCtx.VerifycodeModel.FindOne(l.ctx, req.Email)
	fmt.Println("1111111111111111111111111111111")
	if err != nil {
		return &types.EmailIdentifyResponse{
			Msg: "验证码错误!",
		}, nil
	}
	fmt.Println("2222222222222222222222222222222")
	if verifyCode.Code != req.VerifyCode {
		return &types.EmailIdentifyResponse{
			Msg: "验证码错误!",
		}, nil
	}
	fmt.Println("3333333333333333333333333333333")
	err = l.svcCtx.VerifycodeModel.Delete(l.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println("4444444444444444444444444444444")

	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	newApply := model.Apply{
		UserId:    userId,
		ScholarId: req.ScholarId,
		ApplyType: 1,
		Email:     sql.NullString{String: req.Email, Valid: true},
	}
	fmt.Println("5555555555555555555555555555555")
	_, err = l.svcCtx.ApplyModel.Insert(l.ctx, &newApply)
	if err != nil {
		return nil, err
	}
	fmt.Println("6666666666666666666666666666666")
	return &types.EmailIdentifyResponse{
		Msg: "验证码正确!",
	}, nil
}
