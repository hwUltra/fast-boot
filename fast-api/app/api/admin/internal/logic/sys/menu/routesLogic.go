package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/jwtx"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoutesLogic {
	return &RoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoutesLogic) Routes() (resp *types.RoutesResp, err error) {
	userId := jwtx.GetUid(l.ctx)
	res, err := l.svcCtx.SysRpc.Routes(l.ctx, &sysPb.RoutesReq{Uid: userId, Types: []int64{1, 2, 3}})
	if err != nil {
		return nil, err
	}
	return &types.RoutesResp{List: GetRoutes(res.List, 0)}, nil
}
