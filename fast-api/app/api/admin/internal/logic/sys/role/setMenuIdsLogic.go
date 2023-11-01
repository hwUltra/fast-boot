package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetMenuIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMenuIdsLogic {
	return &SetMenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMenuIdsLogic) SetMenuIds(req *types.SysSetMenuIdsReq) (err error) {
	if _, err = l.svcCtx.SysRpc.RoleSetMenuIds(l.ctx, &sysPb.RoleSetMenuIdsReq{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	}); err != nil {
		return err
	}
	return nil
}
