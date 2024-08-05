package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopListLogic {
	return &ShopListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopListLogic) ShopList(req *types.ShopListReq) (resp *types.ShopListResp, err error) {
	res, err := l.svcCtx.PmsRpc.ShopList(l.ctx, &pmsPb.ListReq{
		Keywords: req.Keywords,
		Status:   req.Status,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.Shop, 0)
	_ = copier.Copy(&list, res.List)
	return &types.ShopListResp{List: list, Total: res.Total}, nil
}
