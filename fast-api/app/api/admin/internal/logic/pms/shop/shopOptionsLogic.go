package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopOptionsLogic {
	return &ShopOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopOptionsLogic) ShopOptions() (resp []*types.OptionItem, err error) {
	options, err := l.svcCtx.PmsRpc.ShopOptions(l.ctx, &pmsPb.AnyReq{})
	if err != nil {
		return nil, err
	}

	list := make([]*types.OptionItem, 0)
	_ = copier.Copy(&list, options.Items)

	return list, nil
}
