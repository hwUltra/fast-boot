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

// 分配菜单权限给角色
func NewSetMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetMenuIdsLogic {
	return &SetMenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetMenuIdsLogic) SetMenuIds(req *types.SysSetMenuIdsReq) (resp *types.NullResp, err error) {
	if _, err := l.svcCtx.SysRpc.RoleSetMenuIds(l.ctx, &sysPb.RoleSetMenuIdsReq{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	}); err != nil {
		return &types.NullResp{}, err
	}
	return &types.NullResp{}, nil
}
