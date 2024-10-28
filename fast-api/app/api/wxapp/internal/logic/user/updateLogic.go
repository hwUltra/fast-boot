package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/hwUltra/fb-tools/jwtx"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UserInfoRep) (resp *types.UserInfoResq, err error) {

	userUpdateReq := umsPb.UserForm{}
	if err := copier.Copy(&userUpdateReq, req); err != nil {
		return nil, err
	}
	userUpdateReq.Id = jwtx.GetUid(l.ctx)
	_, err = l.svcCtx.UmsRpc.UserUpdate(l.ctx, &userUpdateReq)
	if err != nil {
		return nil, err
	}
	return resp, err
}
