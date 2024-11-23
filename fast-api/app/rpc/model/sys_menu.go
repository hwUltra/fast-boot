package model

import (
	"fast-boot/common/cachex"
	"fmt"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/hwUltra/fb-tools/gormx/types"
	"strings"
)

//type MenuType string
//const (
//	CATALOG MenuType = "CATALOG"
//	MENUS   MenuType = "MENU"
//	BUTTON  MenuType = "BUTTON"
//	EXTLINK MenuType = "EXTLINK"
//)

type SysMenuModel struct {
	gormx.BaseDel
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

// -----------------
// for SysMenuCache CacheFun
// -----------------

const CacheSysMenuModelIdPrefix = "Cache:SysMenuModel:ID:"
const CacheSysMenuRoutesUIdPrefix = "Cache:SysMenuRoutes:UID:"

type SysMenuCache gormx.GormCache

func (m *SysMenuCache) Del(idStr string) {
	ids := strings.Split(idStr, ",")
	m.Db.Delete(&SysMenuModel{}, ids)
	for _, id := range ids {
		cacheKey := fmt.Sprintf("%s%v", CacheSysMenuModelIdPrefix, id)
		_ = m.Cache.Del(cacheKey)
	}
}

func (m *SysMenuCache) Create(u *SysMenuModel) error {
	if err := m.Db.Create(&u).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("%s%v", CacheSysMenuModelIdPrefix, u.Id)
	return m.Cache.Set(cacheKey, u)
}

func (m *SysMenuCache) Update(u *SysMenuModel) error {
	if err := m.Db.Save(&u).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("%s%v", CacheSysMenuModelIdPrefix, u.Id)
	return m.Cache.Set(cacheKey, u)
}

func (m *SysMenuCache) Get(id int64) *SysMenuModel {
	cacheKey := fmt.Sprintf("%s%v", CacheSysMenuModelIdPrefix, id)
	item := SysMenuModel{}
	_ = m.Cache.Take(&item, cacheKey, func(user any) error {
		return m.Db.Where("id = ?", id).Model(SysMenuModel{}).First(&user).Error
	})
	return &item
}

func (m *SysMenuCache) Routes(uid int64) []*SysMenuModel {
	cacheKey := fmt.Sprintf("%s%v", CacheSysMenuRoutesUIdPrefix, uid)
	items := make([]*SysMenuModel, 0)
	_ = m.Cache.Take(&items, cacheKey, func(val any) error {
		roleIds := make([]int64, 0)
		if err := m.Db.Model(&SysUserRoleModel{}).
			Joins("JOIN sys_role on sys_role.id = sys_user_role.role_id").
			Where("sys_user_role.user_id = ?", 1).
			Where("sys_role.status = ?", 1).
			Pluck("sys_role.id", &roleIds).Error; err != nil {
			return err
		}

		menuIds := make([]int64, 0)
		if err := m.Db.Model(&SysRoleMenuModel{}).
			Joins("JOIN sys_menu on sys_menu.id = sys_role_menu.menu_id").
			Where("sys_role_menu.role_id IN ?", roleIds).
			Pluck("sys_role_menu.menu_id", &menuIds).Error; err != nil {
			return err
		}

		if err := m.Db.Model(&SysMenuModel{}).
			Where("id IN ? and type in ?", menuIds, []int64{1, 2, 3}).
			Order("sort asc,id asc").Preload("Roles", "status = ?", 1).Find(&items).Error; err != nil {
			return err
		}
		return nil
	})

	return items
}

func (m *SysMenuCache) RoutesClear() {
	cachex.NewStore(m.Conf).ClearRedisPrefix(CacheSysMenuRoutesUIdPrefix)
}
