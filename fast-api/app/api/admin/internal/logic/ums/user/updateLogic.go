package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
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

// 修改用户信息
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UserForm) (resp *types.UserFormResp, err error) {
	form := umsPb.UserForm{}
	if err := copier.Copy(&form, req); err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.UmsRpc.UserUpdate(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.UserFormResp{}, nil
}
