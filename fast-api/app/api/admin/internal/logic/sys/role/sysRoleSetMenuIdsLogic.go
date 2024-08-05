package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleSetMenuIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleSetMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleSetMenuIdsLogic {
	return &SysRoleSetMenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleSetMenuIdsLogic) SysRoleSetMenuIds(req *types.SysSetMenuIdsReq) error {
	if _, err := l.svcCtx.SysRpc.RoleSetMenuIds(l.ctx, &sysPb.RoleSetMenuIdsReq{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	}); err != nil {
		return err
	}
	return nil
}
