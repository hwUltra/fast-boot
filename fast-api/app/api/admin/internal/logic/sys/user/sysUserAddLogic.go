package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserAddLogic {
	return &SysUserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserAddLogic) SysUserAdd(req *types.SysUserFormReq) (resp *types.SysUserAddResp, err error) {
	userForm := &sysPb.UserAddReq{}
	if err := copier.Copy(&userForm, req); err != nil {
		return nil, err
	}
	res, err := l.svcCtx.SysRpc.UserAdd(l.ctx, userForm)
	if err != nil {
		return nil, err
	}
	return &types.SysUserAddResp{
		Id: res.Id,
	}, nil
}
