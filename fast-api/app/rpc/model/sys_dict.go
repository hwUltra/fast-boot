package model

import (
	"fast-boot/common/gormV2"
)

type SysDictModel struct {
	gormV2.BaseDel
	TypeCode string `gorm:"column:type_code;not null" json:"type_code"`     // 字典类型编码
	Name     string `gorm:"column:name;not null" json:"name"`               // 字典项名称
	Value    string `gorm:"column:value;not null" json:"value"`             // 字典项值
	Sort     int64  `gorm:"column:sort;not null" json:"sort"`               // 排序
	Status   int8   `gorm:"column:status;not null;default:1" json:"status"` // 状态(1:正常;0:禁用)
	Remark   string `gorm:"column:remark;not null" json:"remark"`           // 备注
}

func (*SysDictModel) TableName() string {
	return "sys_dict"
}
