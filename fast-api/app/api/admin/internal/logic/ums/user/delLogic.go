package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户
func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.UserDelReq) (resp *types.UserDelResp, err error) {
	if _, err := l.svcCtx.UmsRpc.UserDel(l.ctx, &umsPb.IdsReq{Ids: req.Ids}); err != nil {
		return nil, err
	}
	return &types.UserDelResp{}, nil
}
