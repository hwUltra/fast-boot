package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateLogic {
	return &CategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryUpdateLogic) CategoryUpdate(req *types.PmsCategoryForm) error {
	form := pmsPb.PmsCategoryForm{}
	_ = copier.Copy(&form, req)

	if _, err := l.svcCtx.PmsRpc.PmsCategoryUpdate(l.ctx, &form); err != nil {
		return err
	}

	return nil
}
