package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopUpdateLogic {
	return &ShopUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopUpdateLogic) ShopUpdate(req *types.ShopForm) error {
	form := pmsPb.ShopForm{}
	_ = copier.Copy(&form, req)

	if _, err := l.svcCtx.PmsRpc.ShopUpdate(l.ctx, &form); err != nil {
		return err
	}

	return nil
}
