package logic

import (
	"context"

	"fast-boot/app/rpc/oms/internal/svc"
	"fast-boot/app/rpc/oms/omsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderGetLogic {
	return &OrderGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -------
func (l *OrderGetLogic) OrderGet(in *omsPb.IdReq) (*omsPb.Order, error) {
	// todo: add your logic here and delete this line

	return &omsPb.Order{}, nil
}
