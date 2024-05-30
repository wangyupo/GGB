package migration

import (
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
)

// Migrate 表和数据迁移
func Migrate() {
	err := global.DB.AutoMigrate(
		system.SysUser{},
		system.SysMenu{},
		system.SysRole{},
		system.SysRoleMenu{},
		system.SysRoleUser{},
		system.SysDictCategory{},
		system.SysDictData{},
	)
	if err != nil {
		fmt.Print(err)
	}
	initSystemData()
}

// 初始化admin账号和菜单
func initSystemData() {
	var adminUser system.SysUser
	global.DB.Where("user_name = ?", "admin").First(&adminUser)

	if adminUser.ID == 0 {
		// 将超级管理员默认密码 hash 处理
		var hashPassword = utils.BcryptHash("admin")

		// 创建 admin（超级管理员） 账户
		var adminUser = system.SysUser{
			UserName: "admin",
			NickName: "超级管理员",
			Email:    "",
			Password: hashPassword,
			Status:   1,
		}
		global.DB.Create(&adminUser)

		// 创建 admin（超级管理员） 角色
		var adminRole = system.SysRole{
			RoleName:    "超级管理员",
			RoleCode:    "ADMIN",
			Description: "系统超级管理员角色",
			Status:      1,
		}
		global.DB.Create(&adminRole)

		// 创建系统管理菜单
		global.DB.Create(&[]system.SysMenu{
			{Label: "系统管理", Path: "", Icon: "", ParentId: 0, Sort: 1, Type: 1},
			{Label: "字典管理", Path: "/system/dict", Icon: "", ParentId: 1, Sort: 1, Type: 1},
			{Label: "字典数据", Path: "/system/dict/data", Icon: "", ParentId: 2, Sort: 1, Type: 2},
			{Label: "菜单管理", Path: "/system/menu", Icon: "", ParentId: 1, Sort: 2, Type: 1},
			{Label: "用户管理", Path: "/system/user", Icon: "", ParentId: 1, Sort: 3, Type: 1},
			{Label: "角色管理", Path: "/system/role", Icon: "", ParentId: 1, Sort: 4, Type: 1},
			{Label: "分配用户", Path: "/system/role/user", Icon: "", ParentId: 6, Sort: 1, Type: 2},
			{Label: "登录日志", Path: "/system/loginLog", Icon: "", ParentId: 1, Sort: 5, Type: 1},
		})

		// 关联 admin（超级管理员） 用户和角色
		global.DB.Create(&system.SysRoleUser{
			UserID: adminUser.ID,
			RoleID: adminRole.ID,
		})

		// 关联 admin（超级管理员） 角色和菜单
		var allMenu []system.SysMenu
		global.DB.Find(&allMenu)
		var adminRoleMenus []system.SysRoleMenu
		for _, menu := range allMenu {
			adminRoleMenus = append(adminRoleMenus, system.SysRoleMenu{
				RoleID: adminRole.ID,
				MenuID: menu.ID,
			})
		}
		global.DB.Create(&adminRoleMenus)
	}
}
