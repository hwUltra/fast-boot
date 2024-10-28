package model

import (
	"github.com/hwUltra/fb-tools/gormV2"
	"gorm.io/gorm"
	"time"
)

type ModuleType string

const (
	LOGIN ModuleType = "LOGIN"
	USER  ModuleType = "USER"
	ROLE  ModuleType = "ROLE"
	DEPT  ModuleType = "DEPT"
	MENU  ModuleType = "MENU"
	DICT  ModuleType = "DICT"
	OTHER ModuleType = "OTHER"
)

// SysLogModel mapped from table <sys_log>
type SysLogModel struct {
	gormV2.Base
	Module         ModuleType     `gorm:"column:module;type:enum('LOGIN','USER','ROLE','DEPT','MENU','DICT','OTHER')" json:"module"` // 日志模块
	Content        string         `gorm:"column:content;not null" json:"content"`                                                    // 日志内容
	RequestURI     string         `gorm:"column:request_uri;not null" json:"request_uri"`                                            // 请求路径
	IP             string         `gorm:"column:ip;not null" json:"ip"`                                                              // IP地址
	Province       string         `gorm:"column:province;not null" json:"province"`                                                  // 省份
	City           string         `gorm:"column:city;not null" json:"city"`                                                          // 城市
	ExecutionTime  int64          `gorm:"column:execution_time;not null" json:"execution_time"`                                      // 执行时间(ms)
	Browser        string         `gorm:"column:browser;not null" json:"browser"`                                                    // 浏览器
	BrowserVersion string         `gorm:"column:browser_version;not null" json:"browser_version"`                                    // 浏览器版本
	Os             string         `gorm:"column:os;not null" json:"os"`                                                              // 终端系统
	CreateBy       int64          `gorm:"column:create_by;not null" json:"create_by"`                                                // 创建人ID
	CreatedAt      time.Time      `gorm:"column:created_at;not null" json:"created_at"`                                              // 创建时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;not null" json:"deleted_at"`                                              // 逻辑删除标识(1-已删除 0-未删除)
}

func (*SysLogModel) TableName() string {
	return "sys_log"
}
