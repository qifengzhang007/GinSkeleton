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
	SignKey          string = "GinSkeleton"
)

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 设置SignKey（类似秘钥）
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// --------------------  JWT   ----------------正式阶段  ↓

// 使用工厂创建一个 JWT 结构体
func CreateMyJWT(SignKey string) *Jwt_Sign {
	if len(SignKey) > 0 {
		SetSignKey(SignKey)
	}
	return &Jwt_Sign{
		[]byte(GetSignKey()),
	}
}

// 定义一个 JWT验签 结构体
type Jwt_Sign struct {
	SigningKey []byte
}

// CreateToken 生成一个token
func (j *Jwt_Sign) CreateToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	token_partA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return token_partA.SignedString(j.SigningKey)
}

// 解析Token
func (j *Jwt_Sign) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 如果 TokenExpired ,只是过期（格式都正确），我们认为他是有效的，接下可以允许刷新操作
				token.Valid = true
				goto label_here
			} else {
				return nil, TokenInvalid
			}
		}
	}
label_here:
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, TokenInvalid
	}
}

// 更新token
func (j *Jwt_Sign) RefreshToken(tokenString string, extraAddSeconds int64) (string, error) {

	if CustomClaims, err := j.ParseToken(tokenString); err == nil {
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return j.CreateToken(*CustomClaims)
	} else {
		return "", err
	}
}
