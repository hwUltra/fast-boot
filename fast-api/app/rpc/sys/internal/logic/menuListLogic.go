package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuList 菜单列表
func (l *MenuListLogic) MenuList(in *sysPb.MenuListReq) (*sysPb.MenuListResp, error) {
	items := make([]*model.SysMenuModel, 0)
	l.svcCtx.GormClient.GormDb.Model(&model.SysMenuModel{}).
		Order("id asc").Find(&items)

	list := make([]*sysPb.SysMenuItem, 0)
	for _, item := range items {
		var menu sysPb.SysMenuItem
		_ = copier.Copy(&menu, item)
		list = append(list, &menu)
	}

	return &sysPb.MenuListResp{List: list}, nil
}
