package model

import "github.com/hwUltra/fb-tools/gormx"

type PmsCategoryAttributeModel struct {
	gormx.Base
	CategoryId int64  `gorm:"column:category_id;not null" json:"categoryId"` // 分类ID
	Name       string `gorm:"column:name;not null" json:"name"`              // 属性名称
	Type       int8   `gorm:"column:type;not null" json:"type"`              // 类型(1:规格;2:属性;)
}

func (*PmsCategoryAttributeModel) TableName() string {
	return "pms_category_attribute"
}
