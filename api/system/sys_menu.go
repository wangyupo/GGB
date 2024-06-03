package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
)

// GetMenuList 列表
func GetMenuList(c *gin.Context) {
	// 获取分页参数
	pageNumber, pageSize := utils.GetPaginationParams(c)
	// 获取其它查询参数
	name := c.Query("name")

	// 声明 system.SysMenu 类型的变量以存储查询结果
	menuList := make([]system.SysMenu, 0)
	var total int64

	// 准备数据库查询
	db := global.DB.Model(&system.SysMenu{})
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
	if err := db.Find(&menuList).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  menuList,
		Total: total,
	}, c)
}

// CreateMenu 新建
func CreateMenu(c *gin.Context) {
	// 声明 system.SysMenu 类型的变量以存储 JSON 数据
	var menu system.SysMenu

	// 绑定 JSON 请求体中的数据到 menu 结构体
	if err := c.ShouldBindJSON(&menu); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 创建 menu 记录
	if err := global.DB.Create(&menu).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to create menu", c)
}

// GetMenu 详情
func GetMenu(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysMenu 类型的变量以存储查询结果
	var menu system.SysMenu

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&menu, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(menu, c)
}

// UpdateMenu 编辑
func UpdateMenu(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysMenu 类型的变量以存储查询结果
	var menu system.SysMenu

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&menu, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&menu); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 更新用户记录
	if err := global.DB.Save(&menu).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to update menu", c)
}

// DeleteMenu 删除
func DeleteMenu(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.DB.Delete(&system.SysMenu{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to deleted menu", c)
}
