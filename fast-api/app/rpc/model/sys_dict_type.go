package model

import (
	"github.com/hwUltra/fb-tools/gormV2"
	"gorm.io/gorm"
)

type SysDictTypeModel struct {
	gormV2.Base
	Name   string `gorm:"column:name;not null" json:"name"`               // 类型名称
	Code   string `gorm:"column:code;not null" json:"code"`               // 类型编码
	Status int8   `gorm:"column:status;not null;default:1" json:"status"` // 状态(1:正常;0:禁用)
	Remark string `gorm:"column:remark;not null" json:"remark"`           // 备注
}

func (*SysDictTypeModel) TableName() string {
	return "sys_dict_type"
}

func (*SysDictTypeModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`name` LIKE ?", "%"+keyword+"%").
				Or("`code` LIKE ?", "%"+keyword+"%")
		}
		return db

	}
}

func (*SysDictTypeModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("`status` = ?", status)
		}
		return db
	}
}
