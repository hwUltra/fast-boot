package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopGetLogic {
	return &ShopGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopGetLogic) ShopGet(req *types.ShopGetReq) (resp *types.Shop, err error) {
	res, err := l.svcCtx.PmsRpc.ShopGet(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.Shop
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}
	return &info, nil
}
