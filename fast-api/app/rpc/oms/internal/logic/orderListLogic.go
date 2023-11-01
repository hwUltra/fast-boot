package logic

import (
	"context"

	"fast-boot/app/rpc/oms/internal/svc"
	"fast-boot/app/rpc/oms/omsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListLogic) OrderList(in *omsPb.ListReq) (*omsPb.OrderListResp, error) {
	// todo: add your logic here and delete this line

	return &omsPb.OrderListResp{}, nil
}
