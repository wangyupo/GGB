package log

import (
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
)

type SysLogOperateService struct{}

// GetSysLogOperateList 获取系统操作日志列表
func (s *SysLogOperateService) GetSysLogOperateList(query request.SysLogOperateQuery, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 system.SysLogOperate 类型的变量以存储查询结果
	sysLogOperateList := make([]system.SysLogOperate, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysLogOperate{})
	if query.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+query.Ip+"%")
	}
	if query.StartDate != "" && query.EndDate != "" {
		db = db.Where("created_at >= ? AND created_at <= ?", query.StartDate, query.EndDate)
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Preload("User").Find(&sysLogOperateList).Error

	// 结果集增加 userName 字段
	results := make([]map[string]interface{}, len(sysLogOperateList))
	for i, operateLog := range sysLogOperateList {
		logMap, _ := utils.ExcludeNestedFields(operateLog, []string{"User"})
		// 添加用户名
		logMap["userName"] = operateLog.User.UserName
		results[i] = logMap
	}

	return results, total, err
}

// CreateSysLogOperate 创建系统操作日志
func (s *SysLogOperateService) CreateSysLogOperate(req system.SysLogOperate) (err error) {
	// 创建 sysLogOperate 记录
	err = global.GGB_DB.Create(&req).Error

	return err
}

// GetSysLogOperate 获取系统操作日志详情
func (s *SysLogOperateService) GetSysLogOperate(sysLogOperateId uint) (sysLogOperate system.SysLogOperate, err error) {
	err = global.GGB_DB.Where("id = ?", sysLogOperateId).First(&sysLogOperate).Error
	return sysLogOperate, err
}

// DeleteSysLogOperate 删除系统操作日志
func (s *SysLogOperateService) DeleteSysLogOperate(sysLogOperateId uint) (err error) {
	err = global.GGB_DB.Where("id = ?", sysLogOperateId).Delete(&system.SysLogOperate{}).Error
	return err
}
