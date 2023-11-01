package model

type SysRoleMenuModel struct {
	RoleID int64 `gorm:"column:role_id;not null" json:"role_id"` // 角色ID
	MenuID int64 `gorm:"column:menu_id;not null" json:"menu_id"` // 菜单ID
}

func (*SysRoleMenuModel) TableName() string {
	return "sys_role_menu"
}
