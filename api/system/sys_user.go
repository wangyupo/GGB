package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/router/system/request"
	sysResp "github.com/wangyupo/GGB/router/system/response"
	"github.com/wangyupo/GGB/utils"
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

	// 声明 system.SysUser 类型的变量以存储找出来的用户
	var systemUser system.SysUser
	// 根据 userName 找用户（userName 是唯一的）
	global.DB.Where("user_name=?", loginForm.UserName).First(&systemUser)
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

	response.SuccessWithDetailed(sysResp.LoginResponse{
		User:      systemUser,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
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
	if err := global.DB.Where("id = ?", userId).First(&systemUser).Error; err != nil {
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
	if err := global.DB.Save(&systemUser).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("密码修改成功", c)
}

// GetSystemUserList 列表
func GetSystemUserList(c *gin.Context) {
	// 获取分页参数
	pageNumber, pageSize := utils.GetPaginationParams(c)
	// 获取其它查询参数
	userName := c.Query("userName")

	// 声明 system.SysUser 类型的变量以存储查询结果
	systemUserList := make([]system.SysUser, 0)
	var total int64

	// 准备数据库查询
	db := global.DB.Model(&system.SysUser{})
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
	db = db.Offset((pageNumber - 1) * pageSize).Limit(pageSize)

	// 执行查询并获取结果
	if err := db.Find(&systemUserList).Error; err != nil {
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

	// 创建 systemUser 记录
	if err := global.DB.Create(&systemUser).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to create systemUser", c)
}

// GetSystemUser 详情
func GetSystemUser(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysUser 类型的变量以存储查询结果
	var systemUser system.SysUser

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&systemUser, id).Error; err != nil {
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
	if err := global.DB.First(&systemUser, id).Error; err != nil {
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

	// 更新用户记录
	if err := global.DB.Save(&systemUser).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to update systemUser", c)
}

// DeleteSystemUser 删除
func DeleteSystemUser(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.DB.Delete(&system.SysUser{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to deleted systemUser", c)
}
