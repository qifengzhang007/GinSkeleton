package token

import (
	"github.com/dgrijalva/jwt-go"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/my_jwt"
	"goskeleton/app/models"
	"time"
)

// 创建 usertoken 工厂

func CreateUserFactory() *userToken {
	return &userToken{
		userJwt: my_jwt.CreateMyJWT(consts.JwtTokenSignKey),
	}
}

type userToken struct {
	userJwt *my_jwt.JwtSign
}

//生成token
func (u *userToken) GenerateToken(userid int64, username string, phone string, expire_at int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customeClaims := my_jwt.CustomClaims{
		UserId: userid,
		Name:   username,
		Phone:  phone,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 10),        // 生效开始时间
			ExpiresAt: int64(time.Now().Unix() + expire_at), // 失效截止时间
		},
	}
	return u.userJwt.CreateToken(customeClaims)
}

// 用户login成功，记录用户token
func (u *userToken) RecordLoginToken(userToken, clientIp string) bool {
	if customeClaims, err := u.userJwt.ParseToken(userToken); err == nil {
		userId := customeClaims.UserId
		expiresAt := customeClaims.ExpiresAt
		return models.CreateUserFactory("").OauthLoginToken(userId, userToken, expiresAt, clientIp)
	} else {
		return false
	}
}

// 刷新token的有效期（默认+3600秒，参见常量配置项）
func (u *userToken) RefreshToken(old_token, client_ip string) (new_token string, res bool) {

	// 解析用户token的数据信息
	_, code := u.isNotExpired(old_token)
	switch code {
	case consts.JwtTokenOK, consts.JwtTokenExpired:
		//如果token已经过期，那么执行更新
		if new_token, err := u.userJwt.RefreshToken(old_token, consts.JwtTokenRefreshExpireAt); err == nil {
			if customeClaims, err := u.userJwt.ParseToken(new_token); err == nil {
				userId := customeClaims.UserId
				expiresAt := customeClaims.ExpiresAt
				if models.CreateUserFactory("").OauthRefreshToken(userId, expiresAt, old_token, new_token, client_ip) {
					return new_token, true
				}
			}
		}
	case consts.JwtTokenInvalid:
		variable.ZapLog.Error(my_errors.ErrorsTokenInvalid)
	}

	return "", false
}

// 销毁token，基本用不到，因为一个网站的用户退出都是直接关闭浏览器窗口，极少有户会点击“注销、退出”等按钮，销毁token其实无多大意义
func (u *userToken) DestroyToken() {

}

// 判断token是否未过期
func (u *userToken) isNotExpired(token string) (*my_jwt.CustomClaims, int) {
	if customeClaims, err := u.userJwt.ParseToken(token); err == nil {

		if time.Now().Unix()-customeClaims.ExpiresAt < 0 {
			// token有效
			return customeClaims, consts.JwtTokenOK
		} else {
			// 过期的token
			return customeClaims, consts.JwtTokenExpired
		}
	} else {
		// 无效的token
		return nil, consts.JwtTokenInvalid
	}
}

// 判断token是否有效（未过期+数据库用户信息正常）
func (u *userToken) IsEffective(token string) bool {
	cutomClaims, code := u.isNotExpired(token)
	if consts.JwtTokenOK == code {
		//if user_item := Model.CreateUserFactory("").ShowOneItem(cutomClaims.UserId); user_item != nil {
		if models.CreateUserFactory("").OauthCheckTokenIsOk(cutomClaims.UserId, token) {
			return true
		}
	}
	return false
}
