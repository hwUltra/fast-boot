package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Brand下拉列表
func NewBrandOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandOptionsLogic {
	return &BrandOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandOptionsLogic) BrandOptions(req *types.BrandGetReq) (resp []*types.BrandOption, err error) {
	options, err := l.svcCtx.PmsRpc.PmsBrandOptions(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	list := make([]*types.BrandOption, 0)
	_ = copier.Copy(&list, options.Items)

	return list, nil
}
