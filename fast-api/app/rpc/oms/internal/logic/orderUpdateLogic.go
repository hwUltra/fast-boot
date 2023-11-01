package logic

import (
	"context"

	"fast-boot/app/rpc/oms/internal/svc"
	"fast-boot/app/rpc/oms/omsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderUpdateLogic) OrderUpdate(in *omsPb.OrderUpdateReq) (*omsPb.SuccessResp, error) {
	// todo: add your logic here and delete this line

	return &omsPb.SuccessResp{}, nil
}
