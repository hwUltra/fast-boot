package callBack

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hwUltra/fb-tools/mq"
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
	payload := map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"}
	client := mq.NewTaskClient(
		mq.ClientConfig{Addr: "192.168.3.88:16379"},
	)

	dispatch, err := client.Dispatch(
		globalkey.AsyncTaskOrder,
		payload,
	)

	fmt.Println("Test_Enqueue = ", dispatch, err)
	return resp, err
}
