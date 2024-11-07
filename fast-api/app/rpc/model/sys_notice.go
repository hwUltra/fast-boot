package model

import (
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/hwUltra/fb-tools/utils"
	"gorm.io/gorm"
	"time"
)

type SysNoticeModel struct {
	gormx.BaseDel
	Title         string    `gorm:"column:title;not null" json:"title"`                     // 通知标题
	Content       string    `gorm:"column:content;not null" json:"content"`                 // 通知内容
	Type          int8      `gorm:"column:type;not null" json:"type"`                       // 通知类型（关联字典编码：notice_type）
	Level         string    `gorm:"column:level;not null" json:"level"`                     // 通知等级（字典code：notice_level）
	TargetType    int8      `gorm:"column:target_type;not null" json:"target_type"`         // 目标类型（1: 全体, 2: 指定）
	TargetUserIds string    `gorm:"column:target_user_ids;not null" json:"target_user_ids"` // 目标人ID集合（多个使用英文逗号,分割）
	PublisherID   int64     `gorm:"column:publisher_id;not null" json:"publisher_id"`       // 发布人ID
	PublishStatus int8      `gorm:"column:publish_status;not null" json:"publish_status"`   // 发布状态（0: 未发布, 1: 已发布, -1: 已撤回）
	PublishTime   time.Time `gorm:"column:publish_time;not null" json:"publish_time"`       // 发布时间
	RevokeTime    time.Time `gorm:"column:revoke_time;not null" json:"revoke_time"`         // 撤回时间
	CreateBy      int64     `gorm:"column:create_by;not null" json:"create_by"`             // 创建人ID
	UpdateBy      int64     `gorm:"column:update_by;not null" json:"update_by"`             // 更新人ID
}

// TableName SysNotice's table name
func (*SysNoticeModel) TableName() string {
	return "sys_notice"
}

func (*SysNoticeModel) WithTitle(title string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(title) > 0 {
			return db.Where("`title` LIKE ?", "%"+title+"%")
		}
		return db
	}
}

func (*SysNoticeModel) WithUid(uid int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if uid > 0 {

			//if SqlType == gormx.PostgresqlType {
			//'1' = ANY(string_to_array(target_user_ids, ','));
			return db.Where("?=ANY(string_to_array(target_user_ids, ','))", utils.ToString(uid))
			//} else {
			//return db.Where("find_in_set(?, `target_user_ids`)", uid)
			//}
		}
		return db
	}
}

func (*SysNoticeModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("`publish_status` = ?", status)
		}
		return db
	}
}

func (*SysNoticeModel) WithCreatedAt(startTime string, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(startTime) > 0 && len(endTime) > 0 {
			return db.Where("`created_at` BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}

func (*SysNoticeModel) WithUpdatedAt(startTime string, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(startTime) > 0 && len(endTime) > 0 {
			return db.Where("`updated_at` BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}
