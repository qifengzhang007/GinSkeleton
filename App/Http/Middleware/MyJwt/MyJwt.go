package MyJwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// --------------------  JWT   ----------------准备阶段  ↓
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "bianche"
)

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// 自定义jwt的声明字段信息+标准字段，参开地址：https://blog.csdn.net/codeSquare/article/details/99288718
type CustomClaims struct {
	ID    int    `form:"userid" json:"userid"`
	Name  string `form:"name" json:"name"`
	Phone string `form:"phone" json:"phone"`
	jwt.StandardClaims
}

// --------------------  JWT   ----------------正式阶段  ↓

// 使用工厂创建一个 JWT 结构体
func CreateMyJWT() *JWT_Sign {
	return &JWT_Sign{
		[]byte(GetSignKey()),
	}
}

// 定义一个 JWT验签 结构体
type JWT_Sign struct {
	SigningKey []byte
}

// CreateToken 生成一个token
func (j *JWT_Sign) CreateToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	token_partA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return token_partA.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT_Sign) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT_Sign) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
