package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserUpdateLogic {
	return &SysUserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserUpdateLogic) SysUserUpdate(req *types.SysUserFormReq) error {
	userForm := sysPb.UserUpdateReq{}
	if err := copier.Copy(&userForm, req); err != nil {
		return err
	}
	_, err := l.svcCtx.SysRpc.UserUpdate(l.ctx, &userForm)
	if err != nil {
		return err
	}

	return nil
}
