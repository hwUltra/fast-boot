package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuDelLogic {
	return &SysMenuDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuDelLogic) SysMenuDel(req *types.SysMenuDelReq) error {
	if _, err := l.svcCtx.SysRpc.MenuDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil
}
