package model

import (
	"fast-boot/common/gormV2"
	"gorm.io/gorm"
)

type SysDeptModel struct {
	gormV2.BaseDel
	Name     string `gorm:"column:name;not null" json:"name"`               // 部门名称
	ParentID int64  `gorm:"column:parent_id;not null" json:"parent_id"`     // 父节点id
	TreePath string `gorm:"column:tree_path;not null" json:"tree_path"`     // 父节点id路径
	Sort     int64  `gorm:"column:sort;not null" json:"sort"`               // 显示顺序
	Status   int8   `gorm:"column:status;not null;default:1" json:"status"` // 状态(1:正常;0:禁用)
	CreateBy int64  `gorm:"column:create_by;not null" json:"create_by"`     // 创建人ID
	UpdateBy int64  `gorm:"column:update_by;not null" json:"update_by"`     // 修改人ID
}

func (*SysDeptModel) TableName() string {
	return "sys_dept"
}

func (*SysDeptModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`name` LIKE ?", "%"+keyword+"%")
		}
		return db

	}
}

func (*SysDeptModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("`status` = ?", status)
		}
		return db
	}
}
