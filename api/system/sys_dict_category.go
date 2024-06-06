package system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
	"gorm.io/gorm"
)

// GetSysDictCategoryList 列表
func GetSysDictCategoryList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	name := c.Query("label")

	// 声明 system.SysDictCategory 类型的变量以存储查询结果
	sysDictCategoryList := make([]system.SysDictCategory, 0)
	var total int64

	// 准备数据库查询
	db := global.DB.Model(&system.SysDictCategory{})
	if name != "" {
		db = db.Where("label LIKE ?", "%"+name+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取分页数据
	err := db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&sysDictCategoryList).Error
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  sysDictCategoryList,
		Total: total,
	}, c)
}

// CreateSysDictCategory 新建
func CreateSysDictCategory(c *gin.Context) {
	// 声明 system.SysDictCategory 类型的变量以存储 JSON 数据
	var req system.SysDictCategory

	// 绑定 JSON 请求体中的数据到 sysDictCategory 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	if !errors.Is(global.DB.Where("label = ?", req.Label).First(&system.SysDictCategory{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典 %s 已存在", req.Label), c)
		return
	}
	if !errors.Is(global.DB.Where("label_code = ?", req.LabelCode).First(&system.SysDictCategory{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典编码 %s 已存在", req.LabelCode), c)
		return
	}

	// 创建 sysDictCategory 记录
	if err := global.DB.Create(&req).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysDictCategory 详情
func GetSysDictCategory(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysDictCategory 类型的变量以存储查询结果
	var sysDictCategory system.SysDictCategory

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&sysDictCategory, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysDictCategory, c)
}

// UpdateSysDictCategory 编辑
func UpdateSysDictCategory(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysDictCategory 类型的变量以存储查询结果
	var req system.SysDictCategory

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&req, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	if !errors.Is(global.DB.Where("label = ? AND id != ?", req.Label, id).First(&system.SysDictCategory{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典 %s 已存在", req.Label), c)
		return
	}
	if !errors.Is(global.DB.Where("label_code = ? AND id != ?", req.LabelCode, id).First(&system.SysDictCategory{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典编码 %s 已存在", req.LabelCode), c)
		return
	}

	// 更新用户记录
	if err := global.DB.Save(&req).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSysDictCategory 删除
func DeleteSysDictCategory(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.DB.Delete(&system.SysDictCategory{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}
