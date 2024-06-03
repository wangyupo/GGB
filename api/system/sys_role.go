package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
)

// GetRoleLIst 列表
func GetRoleLIst(c *gin.Context) {
	// 获取分页参数
	pageNumber, pageSize := utils.GetPaginationParams(c)
	// 获取其它查询参数
	name := c.Query("name")

	// 声明 system.SysRole 类型的变量以存储查询结果
	roleLIst := make([]system.SysRole, 0)
	var total int64

	// 准备数据库查询
	db := global.DB.Model(&system.SysRole{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
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
	if err := db.Find(&roleLIst).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  roleLIst,
		Total: total,
	}, c)
}

// CreateRole 新建
func CreateRole(c *gin.Context) {
	// 声明 system.SysRole 类型的变量以存储 JSON 数据
	var role system.SysRole

	// 绑定 JSON 请求体中的数据到 role 结构体
	if err := c.ShouldBindJSON(&role); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 创建 role 记录
	if err := global.DB.Create(&role).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to create role", c)
}

// GetRole 详情
func GetRole(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysRole 类型的变量以存储查询结果
	var role system.SysRole

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&role, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(role, c)
}

// UpdateRole 编辑
func UpdateRole(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysRole 类型的变量以存储查询结果
	var role system.SysRole

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&role, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&role); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 更新用户记录
	if err := global.DB.Save(&role).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to update role", c)
}

// DeleteRole 删除
func DeleteRole(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.DB.Delete(&system.SysRole{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to deleted role", c)
}
