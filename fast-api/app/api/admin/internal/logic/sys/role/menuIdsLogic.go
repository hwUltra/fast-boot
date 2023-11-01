package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuIdsLogic {
	return &MenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuIdsLogic) MenuIds(req *types.PathIdReq) (resp *types.SysMenuIdsResp, err error) {
	result, err := l.svcCtx.SysRpc.RoleMenuIds(l.ctx, &sysPb.IdReq{
		Id: req.Id,
	})
	items := make([]int64, 0)
	if err != nil {
		return nil, err
	}
	if result.Items != nil {
		items = result.Items
	}

	return &types.SysMenuIdsResp{Items: items}, nil
}
