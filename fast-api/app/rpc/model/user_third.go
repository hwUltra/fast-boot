package model

import (
	"github.com/hwUltra/fb-tools/gormx"
	"gorm.io/gorm"
)

type UserThirdModel struct {
	gormx.BaseDel
	Uid      int64      `gorm:"column:uid;not null" json:"uid"`           // 会员ID
	Platform string     `gorm:"column:platform;not null" json:"platform"` // 第三方应用
	Openid   string     `gorm:"column:openid;not null" json:"openid"`     // 第三方唯一ID
	Unionid  string     `gorm:"column:unionid;not null" json:"unionid"`   // 小程序unionid
	Nickname string     `gorm:"column:nickname;not null" json:"nickname"` // 第三方会员昵称
	Avatar   string     `gorm:"column:avatar;not null" json:"avatar"`     // 第三方会员头像
	User     *UserModel `gorm:"foreignKey:uid"`
}

func (*UserThirdModel) TableName() string {
	return "user_third"
}

func (*UserThirdModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`nickname` LIKE ?", "%"+keyword+"%").
				Or("`openid` LIKE ?", "%"+keyword+"%").
				Or("`unionid` LIKE ?", "%"+keyword+"%")
		}
		return db
	}
}
