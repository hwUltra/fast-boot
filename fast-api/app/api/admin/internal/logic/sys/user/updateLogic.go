package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.SysUserFormReq) (resp *types.SysUserFormResp, err error) {
	userForm := sysPb.UserUpdateReq{}
	if err := copier.Copy(&userForm, req); err != nil {
		return resp, err
	}
	_, err = l.svcCtx.SysRpc.UserUpdate(l.ctx, &userForm)
	if err != nil {
		return resp, err
	}

	return resp, err
}
