package menu

import (
	"fast-boot/app/api/admin/internal/types"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"
)

// GetRoutes 转化树结构
func GetRoutes(list []*sysPb.SysMenu, parentId int64) []*types.RouteData {
	res := make([]*types.RouteData, 0)
	for _, item := range list {
		if item.ParentId == parentId {
			var menu types.RouteData
			menu.Id = item.Id
			menu.Path = item.Path
			menu.Component = item.Component
			menu.Name = item.Name
			menu.Redirect = item.Redirect
			menu.Meta = types.RouteMetaData{
				Title:     item.Name,
				Icon:      item.Icon,
				Roles:     item.Roles,
				Hidden:    !(item.Visible > 0),
				KeepAlive: true,
			}
			menu.Children = GetRoutes(list, item.Id)
			if len(menu.Children) > 0 {
				menu.HasChildren = true
			}
			res = append(res, &menu)
		}
	}
	return res
}

// GetTreeList 转化树结构
func GetTreeList(list []*sysPb.SysMenuItem, parentId int64) []*types.SysMenuItem {
	res := make([]*types.SysMenuItem, 0)
	for _, item := range list {
		if item.ParentId == parentId {
			var menu types.SysMenuItem
			_ = copier.Copy(&menu, item)
			menu.Children = GetTreeList(list, item.Id)
			res = append(res, &menu)
		}
	}
	return res
}
