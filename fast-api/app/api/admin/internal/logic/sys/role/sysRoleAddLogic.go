package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleAddLogic {
	return &SysRoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleAddLogic) SysRoleAdd(req *types.SysRoleFormReq) (resp *types.SysRoleAddResp, err error) {
	res, err := l.svcCtx.SysRpc.RoleAdd(l.ctx, &sysPb.RoleForm{
		Name:      req.Name,
		Code:      req.Code,
		Sort:      req.Sort,
		Status:    req.Status,
		DataScope: req.DataScope,
	})
	if err != nil {
		return nil, err
	}

	return &types.SysRoleAddResp{
		Id: res.Id,
	}, nil
}
