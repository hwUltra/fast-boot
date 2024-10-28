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

// 获取角色的菜单ID集合
func NewMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuIdsLogic {
	return &MenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuIdsLogic) MenuIds(req *types.PathIdReq) (resp []int64, err error) {
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

	return items, nil
}
