package model

import (
	"github.com/hwUltra/fb-tools/gormx"
	"gorm.io/gorm"
)

type PmsGoodsModel struct {
	gormx.BaseDel
	Name        string                   `gorm:"column:name;not null" json:"name"`                // 商品名称
	ShopId      int64                    `gorm:"column:shop_id;not null;default:1" json:"shopId"` // shopID
	CategoryId  int64                    `gorm:"column:category_id;not null" json:"categoryId"`   // 商品类型ID
	BrandId     int64                    `gorm:"column:brand_id;not null" json:"brandId"`         // 商品品牌ID
	OriginPrice int64                    `gorm:"column:origin_price;not null" json:"originPrice"` // 原价【起】
	Price       int64                    `gorm:"column:price;not null" json:"price"`              // 现价【起】
	Sales       int64                    `gorm:"column:sales;not null" json:"sales"`              // 销量
	PicUrl      string                   `gorm:"column:pic_url;not null" json:"picUrl"`           // 商品主图
	Album       string                   `gorm:"column:album;not null" json:"album"`              // 商品图册
	Unit        string                   `gorm:"column:unit;not null" json:"unit"`                // 单位
	Description string                   `gorm:"column:description;not null" json:"description"`  // 商品简介
	Detail      string                   `gorm:"column:detail;not null" json:"detail"`            // 商品详情
	Status      int8                     `gorm:"column:status;not null;default:1" json:"status"`  // 商品状态(0:下架 1:上架)
	Shop        PmsShopModel             `gorm:"foreignKey:ShopId"`
	Category    PmsCategoryModel         `gorm:"foreignKey:CategoryId"`
	Brand       PmsBrandModel            `gorm:"foreignKey:BrandId"`
	SkuList     []PmsSkuModel            `gorm:"foreignKey:GoodsId"`
	AttrList    []PmsGoodsAttributeModel `gorm:"foreignKey:GoodsId"`
	SpecList    []PmsGoodsAttributeModel `gorm:"foreignKey:GoodsId"`
}

func (*PmsGoodsModel) TableName() string {
	return "pms_goods"
}

func (*PmsGoodsModel) WithShopId(shopId int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if shopId > 0 {
			return db.Where("shop_id = ?", shopId)
		}
		return db
	}
}

func (*PmsGoodsModel) WithCategoryId(categoryId int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if categoryId > 0 {
			return db.Where("category_id = ?", categoryId)
		}
		return db
	}
}

func (*PmsGoodsModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("status = ?", status)
		}
		return db
	}
}

func (*PmsGoodsModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("name LIKE ?", "%"+keyword+"%")
		}
		return db
	}
}

func (*PmsGoodsModel) WithCreatedAt(startTime string, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(startTime) > 0 && len(endTime) > 0 {
			return db.Where("created_at BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}

// -----------------
// for PmsGoodsModel CacheFun
// -----------------

const CachePmsGoodsModelIdPrefix = "Cache:PmsGoodsModel:ID:"

type PmsGoodsCache gormx.CacheTool
