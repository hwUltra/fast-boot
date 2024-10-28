package model

import (
	"github.com/hwUltra/fb-tools/gormV2"
	"gorm.io/gorm"
)

type UserModel struct {
	gormV2.BaseDel
	Mobile    string             `gorm:"column:mobile;not null" json:"mobile"`             // 电话号码
	Username  string             `gorm:"column:username;not null" json:"username"`         // 姓名
	Nickname  string             `gorm:"column:nickname;not null" json:"nickname"`         // 昵称
	Avatar    string             `gorm:"column:avatar;not null" json:"avatar"`             // 头像
	IdCard    string             `gorm:"column:id_card;not null" json:"id_card"`           // 身份证号码
	Gender    int8               `gorm:"column:gender;not null" json:"gender"`             // 性别 0 未知 1男 2女
	Signature string             `gorm:"column:signature;not null" json:"signature"`       // 签名
	Birthday  string             `gorm:"column:birthday" json:"birthday"`                  // 生日
	Tags      string             `gorm:"column:tags" json:"tags"`                          // tags
	Source    string             `gorm:"column:source;not null;default:APP" json:"source"` // 来源，APP H5
	SourceUid int64              `gorm:"column:source_uid;not null" json:"source_uid"`     // 邀请uid
	UserThird *UserThirdModel    `gorm:"foreignKey:uid"`
	Addresses []UserAddressModel `gorm:"foreignKey:uid"`
}

func (*UserModel) TableName() string {
	return "user"
}

func (*UserModel) WithCreatedAt(startTime string, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(startTime) > 0 && len(endTime) > 0 {
			return db.Where("`created_at` BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}

func (*UserModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("`username` LIKE ?", "%"+keyword+"%").
				Or("`nickname` LIKE ?", "%"+keyword+"%").
				Or("`mobile` LIKE ?", "%"+keyword+"%")
		}
		return db
	}
}
