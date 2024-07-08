package initialize

import (
	"errors"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
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
	// 1-创建数据表
	db := global.GGB_DB
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysRole{},
		system.SysMenu{},
		system.SysRoleUser{},
		system.SysRoleMenu{},
		system.SysDictCategory{},
		system.SysDictData{},

		log.SysLogLogin{},
		log.SysLogOperate{},

		common.UploadFile{},
		common.Transcript{},
	)
	if err != nil {
		global.GGB_LOG.Error("创建数据表错误", zap.Error(err))
		panic(err)
	}

	// 1-1 创建自定义连接表
	err = db.SetupJoinTable(&system.SysUser{}, "Roles", &system.SysRoleUser{})
	err = db.SetupJoinTable(&system.SysRole{}, "Menus", &system.SysRoleMenu{})
	if err != nil {
		global.GGB_LOG.Error("创建自定义连接表错误", zap.Error(err))
		panic(err)
	}

	// 2-初始化默认数据
	initSystemData()
}

// 初始化admin账号和菜单
func initSystemData() {
	var adminUser system.SysUser
	err := global.GGB_DB.Where("user_name = ?", "admin").First(&adminUser).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 判断admin账户是否已创建
			var adminPassword = "admin"                        // 默认超管密码
			var hashPassword = utils.BcryptHash(adminPassword) // 将超级管理员默认密码 hash 处理

			// 1-创建 admin（超级管理员） 账户、角色、账户和角色的连接表
			adminUser = system.SysUser{
				UserName: "admin",
				NickName: "超级管理员",
				Email:    "",
				Password: hashPassword,
				Status:   1,
				Roles: []*system.SysRole{
					{
						RoleName:    "超级管理员",
						RoleCode:    "ADMIN",
						Description: "系统超级管理员角色",
						Status:      1,
					},
				},
			}
			err = global.GGB_DB.Create(&adminUser).Error
			if err != nil {
				global.GGB_LOG.Error("写入超级用户失败！", zap.Error(err))
				panic(err)
			}

			// 2-创建系统菜单、菜单和角色的连接表
			var sysRole []*system.SysRole
			global.GGB_DB.Find(&sysRole, 1)
			menus := []system.SysMenu{
				// 系统管理
				{Label: "系统管理", Path: "", Icon: "", ParentId: 0, Sort: 1, Type: 1, Roles: sysRole},
				{Label: "用户管理", Path: "/systemManage/user", Icon: "", ParentId: 1, Sort: 1, Type: 1, Roles: sysRole},
				{Label: "角色管理", Path: "/systemManage/role", Icon: "", ParentId: 1, Sort: 2, Type: 1, Roles: sysRole},
				{Label: "分配用户", Path: "/systemManage/role/user", Icon: "", ParentId: 3, Sort: 1, Type: 2, Roles: sysRole},
				{Label: "菜单管理", Path: "/systemManage/menu", Icon: "", ParentId: 1, Sort: 3, Type: 1, Roles: sysRole},
				{Label: "字典管理", Path: "/systemManage/dict", Icon: "", ParentId: 1, Sort: 4, Type: 1, Roles: sysRole},
				{Label: "字典数据", Path: "/systemManage/dict/data", Icon: "", ParentId: 6, Sort: 1, Type: 2, Roles: sysRole},
				// 日志管理
				{Label: "日志管理", Path: "", Icon: "", ParentId: 0, Sort: 2, Type: 1, Roles: sysRole},
				{Label: "登录日志", Path: "/logManage/loginLog", Icon: "", ParentId: 8, Sort: 1, Type: 1, Roles: sysRole},
				{Label: "操作日志", Path: "/logManage/operateLog", Icon: "", ParentId: 8, Sort: 2, Type: 1, Roles: sysRole},
				// 示例页面
				{Label: "示例页面", Path: "", Icon: "", ParentId: 0, Sort: 3, Type: 1, Roles: sysRole},
				{Label: "文件管理", Path: "/demo/file", Icon: "", ParentId: 11, Sort: 1, Type: 1, Roles: sysRole},
				{Label: "Excel导入/导出", Path: "/demo/excel", Icon: "", ParentId: 11, Sort: 2, Type: 1, Roles: sysRole},
			}
			err = global.GGB_DB.Create(&menus).Error
			if err != nil {
				global.GGB_LOG.Error("写入系统菜单失败！", zap.Error(err))
				panic(err)
			}
		} else {
			global.GGB_LOG.Error("数据表原始数据填充错误！", zap.Error(err))
			panic(err)
		}
	}
}
