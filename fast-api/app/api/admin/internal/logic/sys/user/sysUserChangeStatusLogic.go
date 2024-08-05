package user

import (
	"context"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserChangeStatusLogic {
	return &SysUserChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserChangeStatusLogic) SysUserChangeStatus(req *types.SysUserChangeStatusReq) error {
	// todo: add your logic here and delete this line

	return nil
}
