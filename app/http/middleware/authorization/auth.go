package authorization

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/myErrors"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"net/http"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

func CheckAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//  模拟验证token
		V_HeaderParams := HeaderParams{}

		//  推荐使用ShouldbindHeader 方式获取头参数
		context.ShouldBindHeader(&V_HeaderParams)

		if len(V_HeaderParams.Authorization) >= 20 {
			token := strings.Split(V_HeaderParams.Authorization, " ")
			if len(token) == 2 && len(token[1]) >= 20 {
				token_is_effective := userstoken.CreateUserFactory().IsEffective(token[1])
				if token_is_effective {
					context.Next()
				} else {
					response.ReturnJson(context, http.StatusUnauthorized, http.StatusUnauthorized, myErrors.Errors_NoAuthorization, "")
					//暂停执行
					context.Abort()
				}
			}
		} else {
			response.ReturnJson(context, http.StatusUnauthorized, http.StatusUnauthorized, myErrors.Errors_NoAuthorization, "")
			//暂停执行
			context.Abort()
		}

	}
}
