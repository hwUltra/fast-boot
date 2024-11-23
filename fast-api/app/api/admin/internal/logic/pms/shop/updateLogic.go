package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改类型
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.ShopForm) (resp *types.ShopFormResp, err error) {
	form := pmsPb.ShopForm{}
	_ = copier.Copy(&form, req)

	if _, err := l.svcCtx.PmsRpc.ShopUpdate(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.ShopFormResp{}, nil
}
