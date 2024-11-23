package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增类型
func NewCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryAddLogic {
	return &CategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryAddLogic) CategoryAdd(req *types.PmsCategoryForm) (resp *types.PmsCategoryFormResp, err error) {
	form := pmsPb.PmsCategoryForm{}
	_ = copier.Copy(&form, req)
	if _, err = l.svcCtx.PmsRpc.PmsCategoryAdd(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.PmsCategoryFormResp{
		Id: req.Id,
	}, nil
}
