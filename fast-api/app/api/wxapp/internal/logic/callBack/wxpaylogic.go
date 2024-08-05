package callBack

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"fast-boot/common/mq"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type WxPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxPayLogic {
	return WxPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxPayLogic) WxPay(req *http.Request) (resp *types.WxPayCallBackResp, err error) {
	//return nil, errors.Wrapf(xerr.NewErrMsg("该用户已被注册"), "用户已经存在 mobile:%s,err:%v", "xxx", err)
	payload := map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"}
	err = mq.DfSend(l.svcCtx.Config.Redis.Host, "orderTask", payload)

	return resp, err
}
