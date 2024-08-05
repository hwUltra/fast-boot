package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UmsUserGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUmsUserGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UmsUserGetLogic {
	return &UmsUserGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UmsUserGetLogic) UmsUserGet(req *types.UmsUserGetReq) (resp *types.UserInfo, err error) {
	res, err := l.svcCtx.UmsRpc.UserGet(l.ctx, &umsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.UserInfo
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}

	return &info, nil
}
