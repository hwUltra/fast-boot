package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UmsUserDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUmsUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UmsUserDelLogic {
	return &UmsUserDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UmsUserDelLogic) UmsUserDel(req *types.UmsUserDelReq) error {
	if _, err := l.svcCtx.UmsRpc.UserDel(l.ctx, &umsPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil
}
