package menu

import (
	"context"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.SysMenuListReq) (resp *types.SysMenuListResp, err error) {
	res, err := l.svcCtx.SysRpc.MenuList(l.ctx, &sysPb.MenuListReq{Keywords: req.Keywords})
	if err != nil {
		return nil, err
	}
	return &types.SysMenuListResp{List: GetTreeList(res.List, 0)}, nil
}
