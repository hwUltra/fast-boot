package model

type SysUserRoleModel struct {
	UserId int64 `gorm:"column:user_id;primaryKey" json:"userId"` // 用户ID
	RoleId int64 `gorm:"column:role_id;primaryKey" json:"roleId"` // 角色ID
}

func (*SysUserRoleModel) TableName() string {
	return "sys_user_role"
}
