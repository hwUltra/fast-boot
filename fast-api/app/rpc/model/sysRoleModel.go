package model

import (
	"fast-boot/common/gormV2"
	"gorm.io/gorm"
)

type SysRoleModel struct {
	gormV2.BaseDel
	Name      string `gorm:"column:name;not null" json:"name"`               // 角色名称
	Code      string `gorm:"column:code;not null" json:"code"`               // 角色编码
	Sort      int64  `gorm:"column:sort;not null" json:"sort"`               // 显示顺序
	Status    int8   `gorm:"column:status;not null;default:1" json:"status"` // 角色状态(1-正常；0-停用)
	DataScope int8   `gorm:"column:data_scope;not null" json:"data_scope"`   // 数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)
}

func (*SysRoleModel) TableName() string {
	return "sys_role"
}

func (*SysRoleModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`name` LIKE ?", "%"+keyword+"%").
				Or("`code` LIKE ?", "%"+keyword+"%")
		}
		return db

	}
}
