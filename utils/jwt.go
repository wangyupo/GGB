package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/wangyupo/GGB/router/system/request"
	"os"
	"time"
)

// CreateClaims 创建token主体信息
func CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 过期时间 7天
		},
	}
	return claims
}

// CreateToken 生成token
func CreateToken(claims request.CustomClaims) (string, error) {
	signingKey := []byte(os.Getenv("TOKEN_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
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
