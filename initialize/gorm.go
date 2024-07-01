package initialize

import (
	"errors"
	"fmt"
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
		system.SysMenu{},
		system.SysRole{},
		system.SysRoleMenu{},
		system.SysRoleUser{},
		system.SysDictCategory{},
		system.SysDictData{},
		log.SysLogLogin{},
		log.SysLogOperate{},
		common.UploadFile{},
	)
	if err != nil {
		fmt.Print(err)
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
			// 将超级管理员默认密码 hash 处理
			var adminPassword = "admin" // 默认超管密码
			var hashPassword = utils.BcryptHash(adminPassword)

			// 1-创建 admin（超级管理员） 账户
			adminUser = system.SysUser{
				UserName: "admin",
				NickName: "超级管理员",
				Email:    "",
				Password: hashPassword,
				Status:   1,
			}
			err = global.GGB_DB.Create(&adminUser).Error
			if err != nil {
				global.GGB_LOG.Error("写入超级用户失败！", zap.Error(err))
			}

			// 2-创建 admin（超级管理员） 角色
			adminRole := system.SysRole{
				RoleName:    "超级管理员",
				RoleCode:    "ADMIN",
				Description: "系统超级管理员角色",
				Status:      1,
			}
			err = global.GGB_DB.Create(&adminRole).Error
			if err != nil {
				global.GGB_LOG.Error("写入超级用户角色失败！", zap.Error(err))
			}

			// 3-创建系统管理菜单
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
				{Label: "操作日志", Path: "/logManage/operateLog", Icon: "", ParentId: 8, Sort: 2, Type: 1},
				// 示例页面
				{Label: "示例页面", Path: "", Icon: "", ParentId: 0, Sort: 3, Type: 1},
				{Label: "文件管理", Path: "/demo/file", Icon: "", ParentId: 11, Sort: 1, Type: 1},
			}
			err = global.GGB_DB.Create(&menus).Error
			if err != nil {
				global.GGB_LOG.Error("写入系统默认菜单失败！", zap.Error(err))
			}

			// 4-关联 admin（超级管理员） 用户和角色
			err = global.GGB_DB.Create(&system.SysRoleUser{
				UserID: adminUser.ID,
				RoleID: adminRole.ID,
			}).Error
			if err != nil {
				global.GGB_LOG.Error("关联超级用户和角色失败！", zap.Error(err))
			}

			// 5-关联 admin（超级管理员） 角色和菜单
			var adminRoleMenus []system.SysRoleMenu
			for _, menu := range menus {
				adminRoleMenus = append(adminRoleMenus, system.SysRoleMenu{
					RoleID: adminRole.ID,
					MenuID: menu.ID,
				})
			}
			err = global.GGB_DB.Create(&adminRoleMenus).Error
			if err != nil {
				global.GGB_LOG.Error("关联超级用户和菜单失败！", zap.Error(err))
			}
		} else {
			global.GGB_LOG.Error("数据表原始数据填充错误", zap.Error(err))
		}
	}
}
