package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"time"
)

type SysBaseService struct{}

const (
	LoginFailedPrefix        = "login_attempts:" // 登录失败，redis 记录的 key 的前缀
	LoginFailedLimitTime     = 5                 // 登录失败次数限制
	LoginFailedResetDuration = 15                // 登录失败超出限制后，重试需要等待的时间（分钟）
)

// isLoginAllowed 检查用户是否超过了允许的登录尝试次数，并返回相应的错误信息
func isLoginAllowed(userName string) (err error) {
	var errStr = ""                                                           // 错误信息
	key := LoginFailedPrefix + userName                                       // 键：登录尝试次数，前缀 + 用户名
	current, err := global.GGB_REDIS.Incr(context.Background(), key).Result() // 增加登录尝试计数
	if err != nil {
		// 如果 Redis 操作发生错误，记录错误日志并返回错误
		global.GGB_LOG.Error("redis incr error", zap.String("err", err.Error()))
		return err
	}

	if current == 1 {
		// 如果这是第一次登录尝试失败，设置键的过期时间
		global.GGB_REDIS.Expire(context.Background(), key, LoginFailedResetDuration*time.Minute)
	}

	if current >= int64(LoginFailedLimitTime) {
		// 如果登录尝试次数超过限制，修改错误信息，提示账户锁定时间
		errStr = fmt.Sprintf("登录失败次数超出限制，账户锁定%d分钟", LoginFailedResetDuration)
	} else {
		// 如果登录尝试次数未超过限制，修改错误信息，提示剩余的尝试次数
		errStr = fmt.Sprintf("用户名或密码错误，还可重试%d次", LoginFailedLimitTime-int(current))
	}

	// 返回包含详细错误信息的错误
	return errors.New(errStr)
}

// isAccountLocked 检查账户是否处于锁定状态
func isAccountLocked(userName string) (bool, error) {
	redisKey := LoginFailedPrefix + userName
	currentAttempts, err := global.GGB_REDIS.Get(context.Background(), redisKey).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		global.GGB_LOG.Error("redis get error", zap.String("err", err.Error()))
		return false, err
	}
	return currentAttempts >= LoginFailedLimitTime, nil
}

// resetLoginAttempts 重置用户的登录尝试计数器
func resetLoginAttempts(userName string) error {
	key := LoginFailedPrefix + userName
	_, err := global.GGB_REDIS.Del(context.Background(), key).Result()
	return err
}

// Login 登录
func (s *SysBaseService) Login(loginForm request.Login) (systemUser system.SysUser, err error) {
	// 检查用户是否超过登录尝试次数限制
	accountLocked, err := isAccountLocked(loginForm.UserName)
	if err != nil {
		return
	}
	if accountLocked {
		err = isLoginAllowed(loginForm.UserName)
		return
	}

	// 根据 userName 找用户（userName 是唯一的）
	err = global.GGB_DB.Preload("Roles").
		Preload("Roles.Menus").Where("user_name = ?", loginForm.UserName).First(&systemUser).Error
	if err != nil {
		// 如果用户不存在或查询错误，处理登录尝试限制
		err = isLoginAllowed(loginForm.UserName)
		return
	}

	// 核对密码
	ok := utils.BcryptCheck(loginForm.Password, systemUser.Password)
	if !ok {
		// 如果密码不匹配，处理登录尝试限制
		err = isLoginAllowed(loginForm.UserName)
		return
	}

	// 登录成功，重置 redis 登录失败计数器
	if err = resetLoginAttempts(systemUser.UserName); err != nil {
		global.GGB_LOG.Error("redis del error", zap.String("err", err.Error()))
		return
	}

	// 核对用户状态
	if systemUser.Status == 0 {
		return systemUser, errors.New("用户已禁用")
	}

	return
}
