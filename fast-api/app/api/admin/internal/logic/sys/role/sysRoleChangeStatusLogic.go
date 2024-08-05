package role

import (
	"context"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleChangeStatusLogic {
	return &SysRoleChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleChangeStatusLogic) SysRoleChangeStatus(req *types.SysRoleChangeStatusReq) error {
	// todo: add your logic here and delete this line

	return nil
}
