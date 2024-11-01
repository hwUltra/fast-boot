package model

import (
	"github.com/hwUltra/fb-tools/gormV2"
	"github.com/hwUltra/fb-tools/gormV2/types"
)

//type MenuType string
//const (
//	CATALOG MenuType = "CATALOG"
//	MENUS   MenuType = "MENU"
//	BUTTON  MenuType = "BUTTON"
//	EXTLINK MenuType = "EXTLINK"
//)

type SysMenuModel struct {
	gormV2.BaseDel
	ParentID   int64          `gorm:"column:parent_id;not null" json:"parentId"` // 父菜单ID
	TreePath   string         `gorm:"column:tree_path;not null" json:"treePath"` // 父节点ID路径
	Name       string         `gorm:"column:name;not null" json:"name"`
	Type       int8           `gorm:"column:type;not null;default:1" json:"type"`
	RouteName  string         `gorm:"column:route_name;not null" json:"routeName"`         // 路由路径(浏览器地址栏路径)
	RoutePath  string         `gorm:"column:route_path;not null" json:"routePath"`         // 路由路径(浏览器地址栏路径)
	Component  string         `gorm:"column:component;not null" json:"component"`          // 组件路径(vue页面完整路径，省略.vue后缀)
	Perm       string         `gorm:"column:perm;not null" json:"perm"`                    // 权限标识
	Visible    bool           `gorm:"column:visible;not null;default:true" json:"visible"` // 显示状态(1-显示;0-隐藏)
	AlwaysShow bool           `gorm:"column:always_show;default:true" json:"alwaysShow"`   // 显示状态(1-显示;0-隐藏)
	KeepAlive  bool           `gorm:"column:keep_alive;default:false" json:"keepAlive"`
	Sort       int64          `gorm:"column:sort;not null" json:"sort"` // 排序
	Icon       string         `gorm:"column:icon;not null" json:"icon"` // 菜单图标
	Redirect   string         `gorm:"column:redirect;not null" json:"redirect"`
	Params     types.JSON     `gorm:"column:params;type:json" json:"params"`
	Roles      []SysRoleModel `gorm:"many2many:sys_role_menu;foreignKey:id;joinForeignKey:menu_id;joinReferences:role_id"`
}

func (*SysMenuModel) TableName() string {
	return "sys_menu"
}
