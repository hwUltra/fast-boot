package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UmsUserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUmsUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UmsUserAddLogic {
	return &UmsUserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UmsUserAddLogic) UmsUserAdd(req *types.UserForm) (resp *types.UmsUserAddResp, err error) {
	form := umsPb.UserForm{}
	if err = copier.Copy(&form, req); err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.UmsRpc.UserAdd(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.UmsUserAddResp{
		Id: req.Id,
	}, nil
}
