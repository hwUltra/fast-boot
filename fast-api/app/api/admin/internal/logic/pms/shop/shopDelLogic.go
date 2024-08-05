package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopDelLogic {
	return &ShopDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopDelLogic) ShopDel(req *types.ShopDelReq) error {
	if _, err := l.svcCtx.PmsRpc.ShopDel(l.ctx, &pmsPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil
}
