package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuUpdateLogic {
	return &SysMenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuUpdateLogic) SysMenuUpdate(req *types.SysMenuForm) error {
	menuForm := sysPb.MenuForm{}
	_ = copier.Copy(&menuForm, req)

	if _, err := l.svcCtx.SysRpc.MenuUpdate(l.ctx, &menuForm); err != nil {
		return err
	}

	return nil
}
