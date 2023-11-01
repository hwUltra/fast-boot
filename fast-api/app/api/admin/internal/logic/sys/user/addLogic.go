package user

import (
	"context"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

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

func (l *AddLogic) Add(req *types.SysUserFormReq) (resp *types.IdResp, err error) {
	userForm := &sysPb.UserAddReq{}
	if err := copier.Copy(&userForm, req); err != nil {
		return nil, err
	}
	res, err := l.svcCtx.SysRpc.UserAdd(l.ctx, userForm)
	if err != nil {
		return nil, err
	}
	return &types.IdResp{
		ID: res.Id,
	}, nil
}
