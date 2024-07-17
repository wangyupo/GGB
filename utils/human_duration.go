package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ParseDuration 解析类似于 "1d7h10m" 的持续时间字符串，并返回 time.Duration 类型的时间间隔
func ParseDuration(d string) (time.Duration, error) {
	// 去除字符串首尾空格
	d = strings.TrimSpace(d)

	// 尝试使用标准库的 ParseDuration 函数解析
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}

	// 如果字符串包含 "d"，需要特殊处理
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		// 获取 "d" 前面的数字，表示天数
		days, err := strconv.Atoi(d[:index])
		if err != nil {
			return 0, errors.New("无效的天数格式: " + d[:index])
		}
		// 天数不能为负
		if days < 0 {
			return 0, errors.New("天数不能为负: " + d[:index])
		}
		dr = time.Hour * 24 * time.Duration(days)

		// 解析 "d" 后面的部分
		remainder := d[index+1:]
		if remainder != "" {
			ndr, err := time.ParseDuration(remainder)
			if err != nil {
				return 0, errors.New("无效的时间间隔格式: " + remainder)
			}
			// 确保时间间隔不是负值
			if ndr < 0 {
				return 0, errors.New("时间间隔不能为负: " + remainder)
			}
			dr += ndr
		}
		return dr, nil
	}

	// 如果字符串不包含任何已知时间单位，返回错误
	return 0, errors.New("无效的时间间隔格式: " + d)
}
