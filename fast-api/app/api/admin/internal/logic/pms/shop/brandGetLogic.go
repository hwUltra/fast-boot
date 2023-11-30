package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandGetLogic {
	return &BrandGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandGetLogic) BrandGet(req *types.PathIdReq) (resp *types.PmsBrand, err error) {
	res, err := l.svcCtx.PmsRpc.PmsBrandGet(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.PmsBrand
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}
	return &info, nil
}
