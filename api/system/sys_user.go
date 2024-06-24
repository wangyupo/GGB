package system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	systemResponse "github.com/wangyupo/GGB/model/system/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	DefaultPassword string = "123456"
)

type SysUserApi struct{}

// Login 登录
func (s *SysUserApi) Login(c *gin.Context) {
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
		global.GGB_LOG.Error("登录失败！获取token失败！", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 设置cookie
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))

	// 写入登录日志
	clientIP := c.ClientIP()               // 获取客户端IP
	userAgent := c.GetHeader("User-Agent") // 获取浏览器信息
	loginLog := log.SysLogLogin{
		UserId:    systemUser.ID,
		Type:      1,
		IP:        clientIP,
		UserAgent: userAgent,
	}
	err = global.GGB_DB.Create(&loginLog).Error
	if err != nil {
		global.GGB_LOG.Error("写入登录日志失败！", zap.Error(err))
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
func (s *SysUserApi) Logout(c *gin.Context) {
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
		Where("sys_role_user.user_id = ? AND sys_role_user.deleted_at IS NULL", userId).
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
		Where("sys_role_menu.role_id = ? AND sys_role_menu.deleted_at IS NULL", role.ID).
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
func (s *SysUserApi) ChangePassword(c *gin.Context) {
	var req request.ChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 从token获取用户id
	userId, err := utils.GetUserID(c)
	if err != nil {
		global.GGB_LOG.Error("修改失败！获取userId失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.ChangePassword(userId, req)
	if err != nil {
		global.GGB_LOG.Error("修改用户密码失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("密码修改成功", c)
}

// ResetPassword 重置用户登录密码
func (s *SysUserApi) ResetPassword(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("用户id不能为空", c)
		return
	}

	err := sysUserService.ResetPassword(id, DefaultPassword)
	if err != nil {
		global.GGB_LOG.Error("重置用户密码失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("密码重置成功！", c)
}

// GetSystemUserInfo 根据token获取用户信息
func (s *SysUserApi) GetSystemUserInfo(c *gin.Context) {
	id, err := utils.GetUserID(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	systemUser, err := sysUserService.GetSystemUserInfo(id)
	if err != nil {
		global.GGB_LOG.Error("获取用户信息失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(systemUser, c)
}

// GetSystemUserList 列表
func (s *SysUserApi) GetSystemUserList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.SystemUserList
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := sysUserService.GetSystemUserList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("获取用户列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateSystemUser 新建
func (s *SysUserApi) CreateSystemUser(c *gin.Context) {
	// 声明 system.SysUser 类型的变量以存储 JSON 数据
	var systemUser system.SysUser

	// 绑定 JSON 请求体中的数据到 systemUser 结构体
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysUserService.CreateSystemUser(systemUser, DefaultPassword)
	if err != nil {
		global.GGB_LOG.Error("新建用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSystemUser 详情
func (s *SysUserApi) GetSystemUser(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	systemUser, err := sysUserService.GetSystemUserInfo(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(systemUser, c)
}

// UpdateSystemUser 编辑
func (s *SysUserApi) UpdateSystemUser(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 声明 system.SysUser 类型的变量以存储查询结果
	var systemUser system.SysUser
	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.UpdateSystemUser(systemUser, id)
	if err != nil {
		global.GGB_LOG.Error("编辑用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSystemUser 删除用户
func (s *SysUserApi) DeleteSystemUser(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.DeleteSystemUser(id)
	if err != nil {
		global.GGB_LOG.Error("删除用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to deleted systemUser", c)
}

// ChangeSystemUserStatus 修改用户状态
func (s *SysUserApi) ChangeSystemUserStatus(c *gin.Context) {
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var req request.ChangeSystemUserStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.ChangeSystemUserStatus(id, req.Status)
	if err != nil {
		global.GGB_LOG.Error("修改用户状态失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}
