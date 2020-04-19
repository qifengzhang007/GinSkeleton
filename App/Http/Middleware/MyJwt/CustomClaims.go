package MyJwt

import "github.com/dgrijalva/jwt-go"

// 自定义jwt的声明字段信息+标准字段，参开地址：https://blog.csdn.net/codeSquare/article/details/99288718
type CustomClaims struct {
	ID    int    `form:"userid" json:"userid"`
	Name  string `form:"name" json:"name"`
	Phone string `form:"phone" json:"phone"`
	jwt.StandardClaims
}
