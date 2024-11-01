package model

import (
	"github.com/hwUltra/fb-tools/gormV2"
	"gorm.io/gorm"
)

type PmsCategoryModel struct {
	gormV2.BaseDel
	ShopId   int64  `gorm:"column:shop_id;not null" json:"shopId"`            // shop_id
	ParentId int64  `gorm:"column:parent_id;not null" json:"parentId"`        // 父级ID
	Name     string `gorm:"column:name;not null" json:"name"`                 // 商品分类名称
	Icon     string `gorm:"column:icon;not null" json:"icon"`                 // 图标地址
	Sort     int64  `gorm:"column:sort;not null;default:1" json:"sort"`       // 排序
	Visible  bool   `gorm:"column:visible;not null;default:1" json:"visible"` // 显示状态:( 0:隐藏 1:显示)
	//Attributes []PmsCategoryAttributeModel `gorm:"foreignKey:CategoryId"`
}

func (*PmsCategoryModel) TableName() string {
	return "pms_category"
}

func (*PmsCategoryModel) WithVisible(visible int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if visible != -1 {
			return db.Where("`visible` = ?", visible)
		}
		return db
	}
}

func (*PmsCategoryModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`name` LIKE ?", "%"+keyword+"%")
		}
		return db
	}
}

func (*PmsCategoryModel) WithShopId(shopId int64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if shopId > 0 {
			return db.Where("`shop_id` = ?", shopId)
		}
		return db
	}
}
