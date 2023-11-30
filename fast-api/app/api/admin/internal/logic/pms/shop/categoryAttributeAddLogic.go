package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryAttributeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryAttributeAddLogic {
	return &CategoryAttributeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryAttributeAddLogic) CategoryAttributeAdd(req *types.PmsCategoryAttributeForm) error {
	if _, err := l.svcCtx.PmsRpc.PmsCategoryAttributeAdd(l.ctx, &pmsPb.PmsCategoryAttributeAddReq{
		CategoryId: req.CategoryId,
		Type:       req.Type,
		Attributes: req.Attributes,
	}); err != nil {
		return err
	}
	return nil
}
