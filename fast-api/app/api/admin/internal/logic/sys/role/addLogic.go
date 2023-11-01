package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.SysRoleFormReq) (resp *types.IdResp, err error) {
	if _, err = l.svcCtx.SysRpc.RoleAdd(l.ctx, &sysPb.RoleForm{
		Name:      req.Name,
		Code:      req.Code,
		Sort:      req.Sort,
		Status:    req.Status,
		DataScope: req.DataScope,
	}); err != nil {
		return nil, err
	}

	return &types.IdResp{
		ID: req.Id,
	}, nil

}
