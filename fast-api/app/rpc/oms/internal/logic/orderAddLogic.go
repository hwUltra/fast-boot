package logic

import (
	"context"

	"fast-boot/app/rpc/oms/internal/svc"
	"fast-boot/app/rpc/oms/omsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderAddLogic) OrderAdd(in *omsPb.OrderAddReq) (*omsPb.SuccessIdResp, error) {
	// todo: add your logic here and delete this line

	return &omsPb.SuccessIdResp{}, nil
}
