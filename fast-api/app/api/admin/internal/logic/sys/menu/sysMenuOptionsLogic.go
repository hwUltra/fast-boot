package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuOptionsLogic {
	return &SysMenuOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuOptionsLogic) SysMenuOptions() (resp []*types.SysMenuOption, err error) {
	options, err := l.svcCtx.SysRpc.MenuOptions(l.ctx, &sysPb.AnyReq{})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysMenuOption, 0)
	_ = copier.Copy(&list, options.Items)

	return list, nil
}
