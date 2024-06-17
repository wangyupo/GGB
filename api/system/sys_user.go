package system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	systemResponse "github.com/wangyupo/GGB/model/system/response"
	"github.com/wangyupo/GGB/utils"
	"gorm.io/gorm"
	"time"
)

// Login 登录
func Login(c *gin.Context) {
	// 声明 loginForm 类型的变量以存储 JSON 数据
	var loginForm request.Login
	if err := c.BindJSON(&loginForm); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	if loginForm.UserName == "" {
		response.FailWithMessage("缺少账户", c)
		return
	}
	if loginForm.Password == "" {
		response.FailWithMessage("缺少密码", c)
		return
	}

	// 声明 system.SysUser 类型的变量以存储找出来的用户
	var systemUser system.SysUser
	// 根据 userName 找用户（userName 是唯一的）
	global.GGB_DB.Where("user_name=?", loginForm.UserName).First(&systemUser)
	if systemUser.ID == 0 {
		response.FailWithMessage("用户不存在", c)
		return
	}
	if systemUser.Status == 2 {
		response.FailWithMessage("用户已禁用", c)
		return
	}

	// 核对密码
	ok := utils.BcryptCheck(loginForm.Password, systemUser.Password)
	if !ok {
		response.FailWithMessage("登录密码错误", c)
		return
	}

	// 密码核对正确，查询用户角色和菜单
	role, menu, err := getRoleMenu(systemUser.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 通过jwt生成token
	claims := utils.CreateClaims(request.BaseClaims{
		ID:       systemUser.ID,
		UserName: systemUser.UserName,
		NickName: systemUser.NickName,
	})
	token, err := utils.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 设置cookie
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))

	// 写入登录日志
	clientIP := c.ClientIP()               // 获取客户端IP
	userAgent := c.GetHeader("User-Agent") // 获取浏览器信息
	loginLog := system.SysLogLogin{
		UserId:    systemUser.ID,
		Type:      1,
		IP:        clientIP,
		UserAgent: userAgent,
	}
	err = global.GGB_DB.Create(&loginLog).Error
	if err != nil {
	}

	response.SuccessWithDetailed(systemResponse.LoginResponse{
		User:      systemUser,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		Role:      role,
		Menu:      menu,
	}, "登录成功", c)
}

// Logout 登出
func Logout(c *gin.Context) {
	utils.ClearToken(c)
	response.SuccessWithDefaultMessage(c)
}

// 获取用户角色和菜单
func getRoleMenu(userId uint) (systemResponse.Role, []systemResponse.Menu, error) {
	var role systemResponse.Role
	var menus []systemResponse.Menu

	// 查找用户角色
	roleErr := global.GGB_DB.Model(&system.SysRole{}).
		Joins("join sys_role_user on sys_role_user.role_id=sys_role.id").
		Where("sys_role_user.user_id = ?", userId).
		First(&role).Error
	if roleErr != nil {
		if errors.Is(roleErr, gorm.ErrRecordNotFound) {
			return systemResponse.Role{}, []systemResponse.Menu{}, fmt.Errorf("账户未匹配角色")
		}
		return systemResponse.Role{}, []systemResponse.Menu{}, roleErr
	}

	// 查找角色对应菜单
	menuErr := global.GGB_DB.Model(&system.SysMenu{}).
		Joins("join sys_role_menu on sys_role_menu.menu_id = sys_menu.id").
		Where("sys_role_menu.role_id = ?", role.ID).
		Scan(&menus).Error
	if menuErr != nil {
		if errors.Is(roleErr, gorm.ErrRecordNotFound) {
			return systemResponse.Role{}, []systemResponse.Menu{}, fmt.Errorf("角色未匹配菜单")
		}
		return systemResponse.Role{}, []systemResponse.Menu{}, menuErr
	}

	return role, menus, nil
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var req request.ChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从token获取用户id
	userId, err := utils.GetUserID(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 根据id查找用户
	var systemUser system.SysUser
	if err := global.GGB_DB.Where("id = ?", userId).First(&systemUser).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验原密码
	if ok := utils.BcryptCheck(req.Password, systemUser.Password); !ok {
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	// 原密码正确，hash新密码，更新sys_user
	systemUser.Password = utils.BcryptHash(req.NewPassword)
	if err := global.GGB_DB.Save(&systemUser).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("密码修改成功", c)
}

// ResetPassword 重置用户登录密码
func ResetPassword(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("用户id不能为空", c)
		return
	}

	err := global.GGB_DB.Model(&system.SysUser{}).
		Where("id = ?", id).
		Update("password", utils.BcryptHash(global.DefaultLoginPassword)).Error
	if err != nil {
		fmt.Print(err)
		response.FailWithMessage("密码重置失败！", c)
		return
	}

	response.SuccessWithMessage("密码重置成功！", c)
}

// GetSystemUserInfo 根据token获取用户信息
func GetSystemUserInfo(c *gin.Context) {
	id, err := utils.GetUserID(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var systemUser system.SysUser
	if err := global.GGB_DB.Model(&system.SysUser{}).Where("id = ?", id).First(&systemUser).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(systemUser, c)
}

// GetSystemUserList 列表
func GetSystemUserList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	userName := c.Query("userName")

	// 声明 system.SysUser 类型的变量以存储查询结果
	systemUserList := make([]system.SysUser, 0)
	var total int64

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysUser{})
	if userName != "" {
		db = db.Where("user_name LIKE ?", "%"+userName+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取分页数据
	err := db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&systemUserList).Error
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  systemUserList,
		Total: total,
	}, c)
}

// CreateSystemUser 新建
func CreateSystemUser(c *gin.Context) {
	// 声明 system.SysUser 类型的变量以存储 JSON 数据
	var systemUser system.SysUser

	// 绑定 JSON 请求体中的数据到 systemUser 结构体
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 检查 UserName 是否重复
	if !errors.Is(global.GGB_DB.Where("user_name = ?", systemUser.UserName).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("用户名 %s 已存在", systemUser.UserName), c)
		return
	}

	systemUser.Password = utils.BcryptHash(global.DefaultLoginPassword)

	// 创建 systemUser 记录
	if err := global.GGB_DB.Create(&systemUser).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSystemUser 详情
func GetSystemUser(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysUser 类型的变量以存储查询结果
	var systemUser system.SysUser

	// 从数据库中查找具有指定 ID 的数据
	if err := global.GGB_DB.First(&systemUser, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(systemUser, c)
}

// UpdateSystemUser 编辑
func UpdateSystemUser(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysUser 类型的变量以存储查询结果
	var systemUser system.SysUser

	// 从数据库中查找具有指定 ID 的数据
	if err := global.GGB_DB.First(&systemUser, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	if !errors.Is(global.GGB_DB.Where("id != ? AND user_name = ?", id, systemUser.UserName).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("用户名 %s 已存在", systemUser.UserName), c)
		return
	}

	// 更新用户记录
	if err := global.GGB_DB.Save(&systemUser).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSystemUser 删除用户
func DeleteSystemUser(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.GGB_DB.Delete(&system.SysUser{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to deleted systemUser", c)
}

// ChangeSystemUserStatus 修改用户状态
func ChangeSystemUserStatus(c *gin.Context) {
	id := c.Param("id")

	var req request.ChangeSystemUserStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := global.GGB_DB.Model(&system.SysUser{}).
		Where("id = ?", id).
		Update("status", req.Status).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}
