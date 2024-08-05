package other

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"fast-boot/common/sms"
	"fast-boot/common/xerr"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendSmsLogic {
	return SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsLogic) SendSms(req types.SendSmsReq) (err error) {
	fmt.Println("sendSms req: ", req)

	var aliSms = sms.NewAliSms(l.svcCtx.Config.Sms.AliConf)
	err = aliSms.SendCode(l.svcCtx.Config.Sms.Template.Reg, req.Mobile, "123213")

	if err != nil {
		return errors.Wrapf(xerr.NewErrMsg(err.Error()), "sendSms fail req: %+v , err : %v ", req, err)

	}

	return nil
}
