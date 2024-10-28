package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Brand列表
func NewBrandPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandPageLogic {
	return &BrandPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandPageLogic) BrandPage(req *types.BrandPageReq) (resp *types.BrandPageResp, err error) {
	res, err := l.svcCtx.PmsRpc.PmsBrandPage(l.ctx, &pmsPb.PmsBrandPageReq{
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
	return &types.BrandPageResp{List: list, Total: res.Total}, nil
}
