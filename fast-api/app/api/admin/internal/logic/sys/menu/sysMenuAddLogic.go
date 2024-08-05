package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuAddLogic {
	return &SysMenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuAddLogic) SysMenuAdd(req *types.SysMenuForm) (resp *types.SysMenuAddResp, err error) {
	menuForm := sysPb.MenuForm{}
	_ = copier.Copy(&menuForm, req)

	if _, err = l.svcCtx.SysRpc.MenuAdd(l.ctx, &menuForm); err != nil {
		return nil, err
	}

	return &types.SysMenuAddResp{
		Id: req.Id,
	}, nil
}
