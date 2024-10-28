package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加用户
func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.UserForm) (resp *types.IdResp, err error) {
	form := umsPb.UserForm{}
	if err = copier.Copy(&form, req); err != nil {
		return nil, err
	}

	if _, err = l.svcCtx.UmsRpc.UserAdd(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.IdResp{
		Id: req.Id,
	}, nil
}
