package system

// SysRoleUser 系统角色和系统用户的自定义连接表
type SysRoleUser struct {
	SysRoleID uint    `json:"sysRoleId" gorm:"primaryKey;comment:角色ID"`
	SysRole   SysRole `json:"sysRole" gorm:"foreignKey:SysRoleID"`
	SysUserID uint    `json:"sysUserId" gorm:"primaryKey;comment:用户ID"`
	SysUser   SysUser `json:"sysUser" gorm:"foreignKey:SysUserID"`
}
