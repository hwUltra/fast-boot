package model

import "github.com/hwUltra/fb-tools/gormV2"

type PmsGoodsAttributeModel struct {
	gormV2.BaseDel
	GoodsId     int64  `gorm:"column:goods_id;not null" json:"goodsId"`
	AttributeId int64  `gorm:"column:attribute_id;not null" json:"attributeId"` // 属性ID
	Name        string `gorm:"column:name;not null" json:"name"`                // 属性名称
	Value       string `gorm:"column:value;not null" json:"value"`              // 属性值
	Type        int8   `gorm:"column:type;not null" json:"type"`                // 类型(1:规格;2:属性;)
	PicUrl      string `gorm:"column:pic_url;not null" json:"picUrl"`           // 规格图片
	Status      int8   `gorm:"column:status;not null;default:1" json:"status"`
}

func (*PmsGoodsAttributeModel) TableName() string {
	return "pms_goods_attribute"
}
