package shop

import (
	"context"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandUpdateLogic {
	return &BrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandUpdateLogic) BrandUpdate(req *types.BrandForm) error {
	// todo: add your logic here and delete this line

	return nil
}
