package model

import (
	"github.com/hwUltra/fb-tools/gormV2"
	"gorm.io/gorm"
)

type SysDictDataModel struct {
	gormV2.Base
	DictId int64  `gorm:"column:dict_id" json:"dictId"`
	Value  string `gorm:"column:value" json:"value"`
	Label  string `gorm:"column:label;not null" json:"label"`
	Tag    string `gorm:"column:tag;not null" json:"tag"`
	Sort   int64  `gorm:"column:sort" json:"sort"`
	Status int8   `gorm:"column:status;not null;default:1" json:"status"`
	Remark string `gorm:"column:remark;not null" json:"remark"`
}

func (*SysDictDataModel) TableName() string {
	return "sys_dict_data"
}

func (*SysDictDataModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("value LIKE ?", "%"+keyword+"%").
				Or("label LIKE ?", "%"+keyword+"%")
		}
		return db

	}
}

func (*SysDictDataModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("status = ?", status)
		}
		return db
	}
}
