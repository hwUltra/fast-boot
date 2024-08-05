package model

import (
	"fast-boot/common/gormV2"
	"gorm.io/gorm"
)

type PmsShopModel struct {
	gormV2.BaseDel          // id
	Name            string  `gorm:"column:name;not null" json:"name"` // 店铺名称
	Tel             string  `gorm:"column:tel;not null" json:"tel"` // 店铺电话
	Notice          string  `gorm:"column:notice;not null" json:"notice"`
	Status          int8    `gorm:"column:status;not null;default:1" json:"status"`                       // 状态(1:正常;0:禁用)// 公告
	DistributionFee float64 `gorm:"column:distribution_fee;not null;default:0.00" json:"distributionFee"` // 配送费
}

func (*PmsShopModel) TableName() string {
	return "pms_shop"
}

func (*PmsShopModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`name` LIKE ?", "%"+keyword+"%").
				Or("`tel` LIKE ?", "%"+keyword+"%")
		}
		return db
	}
}

func (*PmsShopModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("`status` = ?", status)
		}
		return db
	}
}

func (*PmsShopModel) WithCreatedAt(startTime string, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(startTime) > 0 && len(endTime) > 0 {
			return db.Where("`created_at` BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}

func (t *PmsShopModel) Insert() {

}

func (*PmsShopModel) Update() {

}
