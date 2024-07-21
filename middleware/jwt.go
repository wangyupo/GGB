package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
)

// 处理权鉴错误
func errorAuth(message string, c *gin.Context) {
	response.NoAuth(message, c)
	utils.ClearToken(c)
	c.Abort()
}

// Jwt 登录认证的中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1-从浏览器 cookie 中拿 x-token
		token := utils.GetToken(c)
		if token == "" {
			errorAuth("未登录或非法访问", c)
			return
		}

		// 2-解析 token 包含的信息
		claims, err := utils.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				errorAuth("您的会话已过期，请重新登录", c)
				return
			}
			if errors.Is(err, jwt.ErrTokenMalformed) {
				errorAuth("令牌格式错误：令牌包含无效的段数", c)
				return
			}
			errorAuth(err.Error(), c)
			return
		}

		// 3-检查 redis 中的 token，实现单会话登录
		userName, _ := utils.GetUserName(c)
		userTokenInRedis, err := global.GGB_REDIS.Get(context.Background(), "auth:token:"+userName).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				errorAuth("您的会话已过期，请重新登录", c)
				return
			}
			errorAuth(err.Error(), c)
			return
		}
		if userTokenInRedis != token {
			errorAuth("账号已在其它地方登录", c)
			return
		}

		// 4-在 Context 上下文储存键值对
		c.Set("claims", claims)

		// 5-继续处理接下来的处理方法
		c.Next()
	}
}
