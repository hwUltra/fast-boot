package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleGetLogic {
	return &SysRoleGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleGetLogic) SysRoleGet(req *types.SysRoleGetReq) (resp *types.SysRoleInfo, err error) {
	result, err := l.svcCtx.SysRpc.RoleGet(l.ctx, &sysPb.IdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var roleInfo types.SysRoleInfo
	if err := copier.Copy(&roleInfo, result); err != nil {
		return nil, err
	}
	return &roleInfo, nil
}