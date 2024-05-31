package request

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID       uint
	UserName string
	NickName string
}
