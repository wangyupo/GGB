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

// GetSysDictDataList 列表
func GetSysDictDataList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	categoryId := c.Query("categoryId")
	label := c.Query("label")

	// 声明 system.SysDictData 类型的变量以存储查询结果
	sysDictDataList := make([]system.SysDictData, 0)
	var total int64

	// 准备数据库查询
	db := global.DB.Model(&system.SysDictData{}).Where("category_id = ?", categoryId)
	if label != "" {
		db = db.Where("label LIKE ?", "%"+label+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取分页数据
	err := db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&sysDictDataList).Error
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  sysDictDataList,
		Total: total,
	}, c)
}

// CreateSysDictData 新建
func CreateSysDictData(c *gin.Context) {
	// 声明 system.SysDictData 类型的变量以存储 JSON 数据
	var req system.SysDictData

	// 绑定 JSON 请求体中的数据到 sysDictData 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	if !errors.Is(global.DB.Where("label = ?", req.Label).First(&system.SysDictData{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典键 %s 已存在", req.Label), c)
		return
	}
	if !errors.Is(global.DB.Where("value = ?", req.Value).First(&system.SysDictData{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典值 %s 已存在", req.Value), c)
		return
	}

	// 创建 sysDictData 记录
	if err := global.DB.Create(&req).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysDictData 详情
func GetSysDictData(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysDictData 类型的变量以存储查询结果
	var sysDictData system.SysDictData

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&sysDictData, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysDictData, c)
}

// UpdateSysDictData 编辑
func UpdateSysDictData(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysDictData 类型的变量以存储查询结果
	var req system.SysDictData

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

	if !errors.Is(global.DB.Where("label = ? AND id != ?", req.Label, id).First(&system.SysDictData{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典键 %s 已存在", req.Label), c)
		return
	}
	if !errors.Is(global.DB.Where("value = ? AND id != ?", req.Value, id).First(&system.SysDictData{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("字典值 %s 已存在", req.Value), c)
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

// DeleteSysDictData 删除
func DeleteSysDictData(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.DB.Delete(&system.SysDictData{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}
