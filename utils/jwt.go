package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system/request"
	"time"
)

// CreateClaims 创建token主体信息
func CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	ep, _ := ParseDuration(global.GGB_CONFIG.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 设置过期时间
			Issuer:    global.GGB_CONFIG.JWT.Issuer,              // 签名的发行者
		},
	}
	return claims
}

// CreateToken 生成token
func CreateToken(claims request.CustomClaims) (string, error) {
	signingKey := []byte(global.GGB_CONFIG.JWT.SigningKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 创建jwt对象（签名方法：jwt.SigningMethodHS256, 声明：claims）
	return token.SignedString(signingKey)                      // 使用密钥进行签名，并返回最终token
}

// ParseToken 解析token
func ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GGB_CONFIG.JWT.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, jwt.ErrTokenInvalidClaims

	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}
