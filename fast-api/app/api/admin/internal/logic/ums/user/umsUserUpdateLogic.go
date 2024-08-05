package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UmsUserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUmsUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UmsUserUpdateLogic {
	return &UmsUserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UmsUserUpdateLogic) UmsUserUpdate(req *types.UserForm) error {
	form := umsPb.UserForm{}
	if err := copier.Copy(&form, req); err != nil {
		return err
	}
	if _, err := l.svcCtx.UmsRpc.UserUpdate(l.ctx, &form); err != nil {
		return err
	}

	return nil
}
