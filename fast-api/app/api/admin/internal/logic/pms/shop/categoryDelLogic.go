package shop

import (
	"context"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDelLogic {
	return &CategoryDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryDelLogic) CategoryDel(req *types.CategoryDelReq) error {
	if _, err := l.svcCtx.PmsRpc.PmsCategoryDel(l.ctx, &pmsPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}

	return nil
}
