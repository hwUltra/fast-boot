package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandListLogic) BrandList(req *types.BrandListReq) (resp *types.BrandListResp, err error) {
	res, err := l.svcCtx.PmsRpc.PmsBrandList(l.ctx, &pmsPb.PmsBrandListReq{
		Keywords: req.Keywords,
		ShopId:   req.ShopId,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return nil, err
	}
	list := make([]types.PmsBrand, 0)
	_ = copier.Copy(&list, res.List)
	return &types.BrandListResp{List: list, Total: res.Total}, nil
}
