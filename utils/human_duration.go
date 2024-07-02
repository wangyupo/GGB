package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ParseDuration 解析类似于 "1d7h10m" 的持续时间字符串，并返回 time.Duration 类型的时间间隔
func ParseDuration(durationStr string) (time.Duration, error) {
	// 去除字符串首尾空格
	durationStr = strings.TrimSpace(durationStr)

	// 定义一些用于匹配时间单位的括号规则
	pairs := []struct {
		unit   string
		length time.Duration
	}{
		{"d", time.Hour * 24},
		{"h", time.Hour},
		{"m", time.Minute},
		{"s", time.Second},
		{"ms", time.Millisecond},
		{"us", time.Microsecond},
		{"ns", time.Nanosecond},
	}

	var total time.Duration

	// 遍历每一个已知时间单位
	for _, pair := range pairs {
		// 查找该单位在字符串中的位置
		i := strings.Index(durationStr, pair.unit)
		if i >= 0 {
			// 获取单位前面的数字
			numStr := durationStr[:i]
			// 尝试移除解析后的数字和单位
			durationStr = durationStr[i+len(pair.unit):]
			// 将该数字转换为整数
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return 0, errors.New("无效的时间间隔格式: " + numStr + pair.unit)
			}
			// 计算该单位对应的 time.Duration 并累加到总时间中
			total += pair.length * time.Duration(num)
		}
	}

	if len(durationStr) > 0 {
		return 0, errors.New("包含未能识别的时间单位: " + durationStr)
	}

	return total, nil
}
