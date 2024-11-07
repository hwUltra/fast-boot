package model

import (
	"github.com/hwUltra/fb-tools/gormx"
	"gorm.io/gorm"
)

type SysUserModel struct {
	gormx.BaseDel
	Username string         `gorm:"column:username;not null" json:"username"`       // 用户名
	Nickname string         `gorm:"column:nickname;not null" json:"nickname"`       // 昵称
	Gender   int8           `gorm:"column:gender;not null;default:1" json:"gender"` // 性别((1:男;2:女))
	Password string         `gorm:"column:password;not null" json:"password"`       // 密码
	DeptID   int64          `gorm:"column:dept_id;not null" json:"dept_id"`         // 部门ID
	Avatar   string         `gorm:"column:avatar;not null" json:"avatar"`           // 用户头像
	Mobile   string         `gorm:"column:mobile;not null" json:"mobile"`           // 联系方式
	Status   int8           `gorm:"column:status;not null;default:1" json:"status"` // 用户状态((1:正常;0:禁用))
	Email    string         `gorm:"column:email;not null" json:"email"`             // 用户邮箱
	Dept     *SysDeptModel  `gorm:"foreignKey:DeptID"`
	Roles    []SysRoleModel `gorm:"many2many:sys_user_role;foreignKey:id;joinForeignKey:user_id;joinReferences:role_id"`
}

func (*SysUserModel) TableName() string {
	return "sys_user"
}

func (*SysUserModel) WithCreatedAt(startTime string, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(startTime) > 0 && len(endTime) > 0 {
			return db.Where("created_at BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}

func (*SysUserModel) WithKeywords(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyword) > 0 {
			return db.Where("sys_user.username LIKE ?", "%"+keyword+"%").
				Or("sys_user.nickname LIKE ?", "%"+keyword+"%").
				Or("sys_user.mobile LIKE ?", "%"+keyword+"%")
		}
		return db

	}
}

func (*SysUserModel) WithStatus(status int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != -1 {
			return db.Where("sys_user.status = ?", status)
		}
		return db
	}
}

func (*SysUserModel) WithDeptId(deptId int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if deptId > 0 {
			return db.Where("sys_user.dept_id = ?", deptId)
		}
		return db
	}
}
