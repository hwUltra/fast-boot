package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 菜单列表
func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.SysMenuPageReq) (resp []*types.SysMenuItem, err error) {
	res, err := l.svcCtx.SysRpc.MenuList(l.ctx, &sysPb.MenuListReq{Keywords: req.Keywords})
	if err != nil {
		return nil, err
	}
	return GetTreeList(res.List, 0), nil
}
