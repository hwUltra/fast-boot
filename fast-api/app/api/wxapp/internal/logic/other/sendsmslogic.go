package other

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"fast-boot/common/sms"
	"fast-boot/common/xerr"
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

func (l *SendSmsLogic) SendSms(req types.SendSmsReq) (resp *types.SendSmsResp, err error) {

	var smsConf = l.svcCtx.Config.SmsConf

	var aliSms = sms.NewAliSms(
		smsConf.AliYun.RegionId,
		smsConf.AliYun.AccessKeyID,
		smsConf.AliYun.AccessKeySecret,
		smsConf.AliYun.SignName)
	err = aliSms.SendCode(smsConf.Template.Reg, "18086635700", "123213")

	if err != nil {
		//return nil, errors.Wrapf(xerr.NewErrMsg(err.Error()), "sendSms fail req: %+v , err : %v ", req, err)
		return nil, xerr.NewErrMsg(err.Error())

	}

	return
}
