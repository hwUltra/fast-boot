package model

import (
	"github.com/hwUltra/fb-tools/gormx"
	"gorm.io/gorm"
)

type PmsSkuModel struct {
	gormx.BaseDel        // id
	GoodsId       int64  `gorm:"column:goods_id;not null" json:"goodsId"`         // goods ID
	SkuSn         string `gorm:"column:sku_sn;not null" json:"skuSn"`             // 商品编码
	Name          string `gorm:"column:name;not null" json:"name"`                // 商品名称
	SpecIds       string `gorm:"column:spec_ids;not null" json:"specIds"`         // 商品规格值，以英文逗号(,)分割
	Price         int64  `gorm:"column:price;not null" json:"price"`              // 商品价格(单位：分)
	Stock         int64  `gorm:"column:stock;not null" json:"stock"`              // 库存数量
	LockedStock   int64  `gorm:"column:locked_stock;not null" json:"lockedStock"` // 库存锁定数量
	PicURL        string `gorm:"column:pic_url;not null" json:"pic_url"`          // 商品图片地址
	Status        int8   `gorm:"column:status;not null;default:1" json:"status"`
}

func (*PmsSkuModel) TableName() string {
	return "pms_sku"
}

func (*PmsSkuModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`name` LIKE ?", "%"+keyword+"%").
				Or("`tel` LIKE ?", "%"+keyword+"%")
		}
		return db
	}
}

func (*PmsSkuModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("`status` = ?", status)
		}
		return db
	}
}

func (*PmsSkuModel) WithCreatedAt(startTime string, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(startTime) > 0 && len(endTime) > 0 {
			return db.Where("`created_at` BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}
