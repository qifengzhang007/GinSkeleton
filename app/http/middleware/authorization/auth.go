package authorization

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

// 检查token权限
func CheckTokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//  模拟验证token
		headerParams := HeaderParams{}

		//  推荐使用 ShouldBindHeader 方式获取头参数
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsValidatorBindParamsFail, zap.Error(err))
			context.Abort()
		}

		if len(headerParams.Authorization) >= 20 {
			token := strings.Split(headerParams.Authorization, " ")
			if len(token) == 2 && len(token[1]) >= 20 {
				tokenIsEffective := userstoken.CreateUserFactory().IsEffective(token[1])
				if tokenIsEffective {
					if customeToken, err := userstoken.CreateUserFactory().ParseToken(token[1]); err == nil {
						key := variable.ConfigYml.GetString("Token.BindContextKeyName")
						// token验证通过，同时绑定在请求上下文
						context.Set(key, customeToken)
					}
					context.Next()
				} else {
					response.ErrorTokenAuthFail(context)
				}
			}
		} else {
			response.ErrorTokenAuthFail(context)
		}
	}
}

// casbin检查用户对应的角色权限是否允许访问接口
func CheckCasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		requstUrl := c.Request.URL.Path
		method := c.Request.Method

		// 模拟请求参数转换后的角色（roleId=2）
		role := "2" // 这里模拟某个用户的roleId=2

		// 这里将用户的id解析为所拥有的的角色，判断是否具有某个权限即可
		isPass, err := variable.Enforcer.Enforce(role, requstUrl, method)
		if err != nil {
			response.ErrorCasbinAuthFail(c, err.Error())
			return
		} else if !isPass {
			response.ErrorCasbinAuthFail(c, "")
		} else {
			c.Next()
		}
	}
}
