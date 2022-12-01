package logic

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"soft2_backend/service/apply/model"
	"time"

	"soft2_backend/service/apply/api/internal/svc"
	"soft2_backend/service/apply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailVerifyCodeLogic {
	return &EmailVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailVerifyCodeLogic) EmailVerifyCode(req *types.EmailVerifyCodeRequest) error {
	vEmail := []string{req.Email}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	err := SendEmail(vEmail, code)
	if err != nil {
		return err
	}
	oldCode, err := l.svcCtx.VerifycodeModel.FindOne(l.ctx, req.Email)
	if oldCode != nil {
		err := l.svcCtx.VerifycodeModel.Delete(l.ctx, req.Email)
		if err != nil {
			return err
		}
	}

	newCode := model.Verifycode{
		Email: req.Email,
		Code:  code,
	}
	_, err = l.svcCtx.VerifycodeModel.Insert(l.ctx, &newCode)
	if err != nil {
		return err
	}

	return nil
}

func SendEmail(address []string, vCode string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("DiamondPick <413935740@qq.com>")
	e.To = address
	t := time.Now().Format("2006-01-02 15:04:05")
	//设置文件发送的内容
	content := fmt.Sprintf(`<div>
	    <div>
	        <img src="https://img1.baidu.com/it/u=2394665572,2612107105&amp;fm=253&amp;fmt=auto&amp;app=138&amp;f=JPEG?w=100&amp;h=100" width="100px" height="100px">
	    </div>
	    <div>
	        尊敬的%s，您好！
	    </div>
	    <div style="padding: 8px 40px 8px 50px;">
	        <p>
	            您于 %s 提交的邮箱验证，本次验证码为
	            <u>
	                <strong>
	                    %s
	                </strong>
	            </u>
	            。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。
	        </p>
	    </div>
	    <div>
	        <p>
	            此邮箱为系统邮箱，请勿回复。
	        </p>
	    </div>
	</div>`, address[0], t, vCode)
	e.HTML = []byte(content)
	e.Subject = "验证码"
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "413935740@qq.com", "ukdwwhkaegvpcbch", "smtp.qq.com"))
	return err
}
