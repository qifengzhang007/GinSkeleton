package my_jwt

import "github.com/dgrijalva/jwt-go"

// 自定义jwt的声明字段信息+标准字段，参考地址：https://blog.csdn.net/codeSquare/article/details/99288718
type CustomClaims struct {
	UserId int64  `form:"userid" json:"userid"`
	Name   string `form:"user_name" json:"user_name"`
	Phone  string `form:"phone" json:"phone"`
	jwt.StandardClaims
}
