package callBack

import (
	"context"
	"net/http"

	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *WxPayLogic) WxPay(rw http.ResponseWriter, req *http.Request) (resp *types.WxPayCallBackResp, err error) {

	return
}
