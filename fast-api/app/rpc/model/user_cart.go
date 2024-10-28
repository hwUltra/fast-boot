package model

import "github.com/hwUltra/fb-tools/gormV2"

type ShopCartModel struct {
	gormV2.Base
	ShopId    int64  `gorm:"column:shop_id;not null" json:"shopId"`            // 商户id
	Uid       int64  `gorm:"column:uid;not null" json:"uid"`                   // 用户id
	GoodsId   int64  `gorm:"column:goods_id;not null" json:"goodsId"`          // 商品id
	GoodsName string `gorm:"column:goods_name;not null" json:"goodsName"`      // 商品名称
	Checked   int8   `gorm:"column:checked;not null;default:1" json:"checked"` // 购物车中商品是否选择状态
	PicURL    string `gorm:"column:pic_url;not null" json:"picUrl"`            // 商品图片或者商品货品图片
	Num       int64  `gorm:"column:num;not null" json:"num"`                   // 数量
	SkuID     int64  `gorm:"column:sku_id;not null" json:"skuId"`              // sku id
	Spec      string `gorm:"column:spec;not null" json:"spec"`                 // spec
	SpecStr   string `gorm:"column:spec_str;not null" json:"specStr"`          // spec 描述
}

func (*ShopCartModel) TableName() string {
	return "shop_cart"
}
