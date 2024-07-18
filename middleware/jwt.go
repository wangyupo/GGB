package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
)

// Jwt 登录认证的中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1-从浏览器 cookie 中拿 x-token
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}

		// 2-解析 token 包含的信息
		claims, err := utils.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				response.NoAuth("授权已过期", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.NoAuth(err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		// 3-检查redis是否已存在该用户的token
		userName, _ := utils.GetUserName(c)
		userTokenInRedis, _ := global.GGB_REDIS.Get(context.Background(), "token_"+userName).Result()
		if userTokenInRedis != token {
			utils.ClearToken(c)
			response.NoAuth("账号已在其它地方登录", c)
			c.Abort()
			return
		}

		// 4-在Context上下文储存键值对
		c.Set("claims", claims)

		// 5-继续处理接下来的处理方法
		c.Next()
	}
}
