package request

import "github.com/golang-jwt/jwt/v5"

// CustomClaims claims 是存储在 JWT 中的一组声明，用来传递有关用户或安全上下文的信息
// JWT 由三个部分组成：Header (头部), Payload (负载), 和 Signature (签名)。claims 在 Payload 部分中。
// Header（头部）：包含应用于 token 的签名算法和 token 类型（通常为 "JWT"）。
// Payload（负载）：包含声明，其中包括 ID、颁发者、主题和其他相关声明（claims）。
// Signature（签名）：头部和负载组合形成签名部分，保证数据一致性。
type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID       uint
	UserName string
	NickName string
}
