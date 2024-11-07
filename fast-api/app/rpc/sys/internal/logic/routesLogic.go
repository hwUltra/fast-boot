package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoutesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoutesLogic {
	return &RoutesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoutesLogic) Routes(in *sysPb.RoutesReq) (*sysPb.RoutesResp, error) {
	gormDb := l.svcCtx.GormClient.GormDb
	items := make([]*model.SysMenuModel, 0)
	if in.Uid > 0 {
		roleIds := make([]int64, 0)
		gormDb.Model(&model.SysUserRoleModel{}).
			Joins("JOIN sys_role on sys_role.id = sys_user_role.role_id").
			Where("sys_user_role.user_id = ?", 1).
			Where("sys_role.status = ?", 1).
			Pluck("sys_role.id", &roleIds)

		menuIds := make([]int64, 0)
		gormDb.Model(&model.SysRoleMenuModel{}).
			Joins("JOIN sys_menu on sys_menu.id = sys_role_menu.menu_id").
			Where("sys_role_menu.role_id IN ?", roleIds).
			Pluck("sys_role_menu.menu_id", &menuIds)

		if len(in.Types) > 0 {
			gormDb.Model(&model.SysMenuModel{}).
				Where("id IN ? and type in ?", menuIds, in.Types).
				Order("sort asc,id asc").Preload("Roles", "status = ?", 1).Find(&items)
		} else {
			gormDb.Model(&model.SysMenuModel{}).Where("id IN ?", menuIds).
				Order("sort asc,id asc").Preload("Roles", "status = ?", 1).Find(&items)
		}

	} else {
		if len(in.Types) > 0 {
			gormDb.Model(&model.SysMenuModel{}).
				Where("type in ?", in.Types).
				Order("sort asc,id asc").
				Preload("Roles", "status = ?", 1).Find(&items)
		} else {
			gormDb.Model(&model.SysMenuModel{}).
				Order("sort asc,id asc").
				Preload("Roles", "status = ?", 1).Find(&items)
		}

	}
	list := make([]*sysPb.SysMenu, 0)
	for _, item := range items {
		var menu sysPb.SysMenu
		_ = copier.Copy(&menu, item)

		roles := make([]string, 0)
		for _, v := range item.Roles {
			roles = append(roles, v.Code)
		}
		menu.Roles = roles

		list = append(list, &menu)
	}

	return &sysPb.RoutesResp{List: list}, nil
}
