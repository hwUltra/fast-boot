package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除角色
func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.SysRoleDelReq) (resp *types.SysRoleDelResp, err error) {
	if _, err := l.svcCtx.SysRpc.RoleDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return nil, err
	}
	return &types.SysRoleDelResp{}, nil
}
