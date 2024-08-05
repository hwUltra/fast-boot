package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserChangePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserChangePwdLogic {
	return &SysUserChangePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserChangePwdLogic) SysUserChangePwd(req *types.SysUserChangePwdReq) error {
	_, err := l.svcCtx.SysRpc.UserChangePwd(l.ctx, &sysPb.UserChangePwdReq{
		UserId:   req.UserId,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	return nil
}
