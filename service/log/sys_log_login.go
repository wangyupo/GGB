package log

import (
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/log"
)

type SysLogLoginService struct{}

func (s *SysLogLoginService) GetSysLogLoginList(userId uint, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 log.SysLogLogin 类型的变量以存储查询结果
	loginLogList := make([]log.SysLogLogin, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&log.SysLogLogin{})
	if userId != 0 {
		db = db.Where("user_id = ?", userId)
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").
		Unscoped().Preload("User").Find(&loginLogList).Error
	if err != nil {
		return
	}

	return loginLogList, total, err
}
