package system

// SysRoleMenu 系统角色和系统菜单的自定义连接表
type SysRoleMenu struct {
	SysRoleID uint     `json:"sysRoleId" gorm:"primaryKey;comment:角色ID"`
	SysRole   *SysRole `json:"sysRole" gorm:"foreignKey:SysRoleID"`
	SysMenuID uint     `json:"sysMenuId" gorm:"primaryKey;comment:菜单ID"`
	SysMenu   *SysMenu `json:"sysMenu" gorm:"foreignKey:SysMenuID"`
}
