package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改Brand
func NewBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandUpdateLogic {
	return &BrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandUpdateLogic) BrandUpdate(req *types.BrandForm) (resp *types.NullResp, err error) {
	form := pmsPb.PmsBrandForm{}
	_ = copier.Copy(&form, req)

	if _, err := l.svcCtx.PmsRpc.PmsBrandUpdate(l.ctx, &form); err != nil {
		return &types.NullResp{}, err
	}

	return &types.NullResp{}, nil
}
