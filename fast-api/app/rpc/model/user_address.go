package model

import "fast-boot/common/gormV2"

type UserAddressModel struct {
	gormV2.BaseDel
	ShopId    int64  `gorm:"column:shop_id;not null" json:"shopId"`       // 店铺id
	Uid       int64  `gorm:"column:uid;not null" json:"uid"`              // 用户ID
	Name      string `gorm:"column:name;not null" json:"name"`            // 联系人
	Mobile    string `gorm:"column:mobile;not null" json:"mobile"`        // 手机号
	Area      string `gorm:"column:area;not null" json:"area"`            // 学校地址
	Info      string `gorm:"column:info;not null" json:"info"`            // 详细地址
	IsDefault int8   `gorm:"column:is_default;not null" json:"isDefault"` // 是否默认
}

func (*UserAddressModel) TableName() string {
	return "user_address"
}
