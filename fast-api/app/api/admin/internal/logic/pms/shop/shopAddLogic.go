package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopAddLogic {
	return &ShopAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopAddLogic) ShopAdd(req *types.ShopForm) (resp *types.ShopAddResp, err error) {
	form := pmsPb.ShopForm{}
	_ = copier.Copy(&form, req)
	if _, err = l.svcCtx.PmsRpc.ShopAdd(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.ShopAddResp{
		Id: req.Id,
	}, nil
}
