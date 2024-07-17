package utils

import (
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
		hour, err := strconv.Atoi(d[:index])
		if err != nil {
			return 0, err
		}
		dr = time.Hour * 24 * time.Duration(hour)

		// 解析 "d" 后面的部分
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			// 如果解析 "d" 后面的部分失败，仅返回天数部分
			return dr, nil
		}
		// 返回天数部分和后面部分的总和
		return dr + ndr, nil
	}

	// 如果字符串不包含任何时间单位，尝试将其解析为整数并返回
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
