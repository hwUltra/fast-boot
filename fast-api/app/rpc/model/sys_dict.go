package model

import (
	"github.com/hwUltra/fb-tools/gormx"
	"gorm.io/gorm"
)

type SysDictModel struct {
	gormx.BaseDel
	Code     string             `gorm:"column:code;not null" json:"type_code"`          // 字典类型编码
	Name     string             `gorm:"column:name;not null" json:"name"`               // 字典项名称
	Status   int8               `gorm:"column:status;not null;default:1" json:"status"` // 状态(1:正常;0:禁用)
	Remark   string             `gorm:"column:remark;not null" json:"remark"`           // 备注
	DataList []SysDictDataModel `gorm:"foreignKey:DictId"`
}

func (*SysDictModel) TableName() string {
	return "sys_dict"
}

func (*SysDictModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("status = ?", status)
		}
		return db
	}
}

func (*SysDictModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("name LIKE ?", "%"+keyword+"%").
				Or("code LIKE ?", "%"+keyword+"%")
		}
		return db

	}
}

// -----------------
// for SysDeptModel CacheFun
// -----------------

const CacheSysDictModelIdPrefix = "Cache:SysDictModel:ID:"

type SysDictCache gormx.GormCache
