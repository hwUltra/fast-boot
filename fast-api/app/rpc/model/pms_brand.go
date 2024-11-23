package model

import (
	"github.com/hwUltra/fb-tools/gormx"
	"gorm.io/gorm"
)

type PmsBrandModel struct {
	gormx.BaseDel
	ShopId int64  `gorm:"column:shop_id;not null" json:"shopId"` // shopID
	Name   string `gorm:"column:name;not null" json:"name"`      // 品牌名称
	Logo   string `gorm:"column:logo;not null" json:"logo"`      // LOGO图片
	Sort   int64  `gorm:"column:sort;not null" json:"sort"`      // 排序
}

// TableName PmsBrand's table name
func (*PmsBrandModel) TableName() string {
	return "pms_brand"
}

func (*PmsBrandModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("name LIKE ?", "%"+keyword+"%")
		}
		return db
	}
}

func (*PmsBrandModel) WithShopId(shopId int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if shopId > 0 {
			return db.Where("shop_id = ?", shopId)
		}
		return db
	}
}

// -----------------
// for PmsBrandModel CacheFun
// -----------------

const CachePmsBrandModelIdPrefix = "Cache:PmsBrandModel:ID:"

type PmsBrandCache gormx.GormCache
