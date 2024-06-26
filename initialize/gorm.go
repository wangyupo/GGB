package initialize

import (
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GGB_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	// 1、创建数据表
	db := global.GGB_DB
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysMenu{},
		system.SysRole{},
		system.SysRoleMenu{},
		system.SysRoleUser{},
		system.SysDictCategory{},
		system.SysDictData{},
		log.SysLogLogin{},
		log.SysLogOperate{},
	)
	if err != nil {
		fmt.Print(err)
	}

	// 2、初始化默认数据
	initSystemData()
}

// 初始化admin账号和菜单
func initSystemData() {
	var adminUser system.SysUser
	global.GGB_DB.Where("user_name = ?", "admin").First(&adminUser)

	if adminUser.ID == 0 {
		// 将超级管理员默认密码 hash 处理
		var adminPassword = "admin" // 默认超管密码
		var hashPassword = utils.BcryptHash(adminPassword)

		// 创建 admin（超级管理员） 账户
		adminUser = system.SysUser{
			UserName: "admin",
			NickName: "超级管理员",
			Email:    "",
			Password: hashPassword,
			Status:   1,
		}
		global.GGB_DB.Create(&adminUser)

		// 创建 admin（超级管理员） 角色
		adminRole := system.SysRole{
			RoleName:    "超级管理员",
			RoleCode:    "ADMIN",
			Description: "系统超级管理员角色",
			Status:      1,
		}
		global.GGB_DB.Create(&adminRole)

		// 创建系统管理菜单
		menus := []system.SysMenu{
			// 系统管理
			{Label: "系统管理", Path: "", Icon: "", ParentId: 0, Sort: 1, Type: 1},
			{Label: "用户管理", Path: "/systemManage/user", Icon: "", ParentId: 1, Sort: 1, Type: 1},
			{Label: "角色管理", Path: "/systemManage/role", Icon: "", ParentId: 1, Sort: 2, Type: 1},
			{Label: "分配用户", Path: "/systemManage/role/user", Icon: "", ParentId: 3, Sort: 1, Type: 2},
			{Label: "菜单管理", Path: "/systemManage/menu", Icon: "", ParentId: 1, Sort: 3, Type: 1},
			{Label: "字典管理", Path: "/systemManage/dict", Icon: "", ParentId: 1, Sort: 4, Type: 1},
			{Label: "字典数据", Path: "/systemManage/dict/data", Icon: "", ParentId: 6, Sort: 1, Type: 2},
			// 日志管理
			{Label: "日志管理", Path: "", Icon: "", ParentId: 0, Sort: 2, Type: 1},
			{Label: "登录日志", Path: "/logManage/loginLog", Icon: "", ParentId: 8, Sort: 1, Type: 1},
		}
		global.GGB_DB.Create(&menus)

		// 关联 admin（超级管理员） 用户和角色
		global.GGB_DB.Create(&system.SysRoleUser{
			UserID: adminUser.ID,
			RoleID: adminRole.ID,
		})

		// 关联 admin（超级管理员） 角色和菜单
		var adminRoleMenus []system.SysRoleMenu
		for _, menu := range menus {
			adminRoleMenus = append(adminRoleMenus, system.SysRoleMenu{
				RoleID: adminRole.ID,
				MenuID: menu.ID,
			})
		}
		global.GGB_DB.Create(&adminRoleMenus)
	}
}
