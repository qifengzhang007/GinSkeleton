package Token

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Http/Middleware/MyJwt"
	"GinSkeleton/App/Model"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// 创建 usertoken 工厂

func CreateUserFactory() *userToken {
	return &userToken{
		userJwt: MyJwt.CreateMyJWT(Consts.JwtToken_SignKey),
	}
}

type userToken struct {
	userJwt *MyJwt.Jwt_Sign
}

//生成token
func (u *userToken) GenerateToken(userid int64, username string, phone string, expire_at int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	custome_claims := MyJwt.CustomClaims{
		UserId: userid,
		Name:   username,
		Phone:  phone,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 10),        // 生效开始时间
			ExpiresAt: int64(time.Now().Unix() + expire_at), // 失效截止时间
		},
	}
	return u.userJwt.CreateToken(custome_claims)
}

// 刷新token的有效期（默认+3600秒，参见常量配置项）
func (u *userToken) RefreshToken(token string) (res bool) {

	// 解析用户token的数据信息
	cutomClaims, code := u.isNotExpired(token)
	switch code {
	case Consts.JwtToken_OK:
		// token本身有效，直接犯规true，无需更新
		res = true
	case Consts.JwtToken_Expired:
		//如果token已经过期，那么执行更新
		if new_token, error := u.userJwt.RefreshToken(token, Consts.JwtToken_Refresh_ExpireAt); error == nil {
			//cutomClaims.ID=userid
			if Model.CreateUserFactory().RefreshToken(cutomClaims.UserId, new_token) {
				res = true
			}
		}
	case Consts.JwtToken_Invalid:
		log.Panic(MyErrors.Errors_Token_Invalid)
	}

	return res
}

// 销毁token
func (u *userToken) DestroyToken() {

}

// 判断token是否未过期
func (u *userToken) isNotExpired(token string) (*MyJwt.CustomClaims, int) {
	if custome_claims, err := u.userJwt.ParseToken(token); err == nil {

		if time.Now().Unix()-custome_claims.ExpiresAt < 0 {
			// token有效
			return custome_claims, Consts.JwtToken_OK
		} else {
			// 过期的token
			return custome_claims, Consts.JwtToken_Expired
		}
	}
	// 无效的token
	return nil, Consts.JwtToken_Invalid
}

// 判断token是否有效（未过期+数据库用户信息正常）
func (u *userToken) IsEffective(token string) bool {
	cutomClaims, code := u.isNotExpired(token)
	if Consts.JwtToken_OK == code {
		if user_item := Model.CreateUserFactory().ShowOneItem(cutomClaims.UserId); user_item != nil {
			if user_item.Token == token {
				return true
			}
		}
	}
	return false
}
