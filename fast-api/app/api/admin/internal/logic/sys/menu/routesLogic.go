package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/hwUltra/fb-tools/jwtx"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoutesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 路由列表
func NewRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoutesLogic {
	return &RoutesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoutesLogic) Routes() (resp []*types.RouteData, err error) {
	userId := jwtx.GetUid(l.ctx)
	res, err := l.svcCtx.SysRpc.Routes(l.ctx, &sysPb.RoutesReq{Uid: userId, Types: []int64{1, 2, 3}})
	if err != nil {
		return nil, err
	}
	//[]*SysMenu
	return GetRoutes(res.List, 0), nil
}
