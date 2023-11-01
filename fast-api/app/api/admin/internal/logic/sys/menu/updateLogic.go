package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
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

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.SysMenuForm) (resp *types.IdResp, err error) {
	menuForm := sysPb.MenuForm{}
	_ = copier.Copy(&menuForm, req)

	if _, err = l.svcCtx.SysRpc.MenuUpdate(l.ctx, &menuForm); err != nil {
		return nil, err
	}

	return &types.IdResp{
		ID: req.Id,
	}, nil
}
