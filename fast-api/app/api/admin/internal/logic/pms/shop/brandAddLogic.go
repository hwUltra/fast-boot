package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandAddLogic {
	return &BrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandAddLogic) BrandAdd(req *types.BrandForm) (resp *types.BrandAddResp, err error) {
	form := pmsPb.PmsBrandForm{}
	_ = copier.Copy(&form, req)
	if _, err = l.svcCtx.PmsRpc.PmsBrandAdd(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.BrandAddResp{
		Id: req.Id,
	}, nil
}
