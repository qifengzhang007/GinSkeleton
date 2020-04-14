package Authorization

import (
	"GinSkeleton/App/Global/Errors"
	"fmt"
	"net/http"
	"strings"

	//"GinSkeleton/Vendors/github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

		if len(V_HeaderParams.Authorization) < 1 {
			context.JSON(401,
				gin.H{
					"code": http.StatusUnauthorized,
					"msg":  Errors.Errors_NoAuthorization,
				})
			//暂停执行
			context.Abort()
		} else {
			token := strings.Split(V_HeaderParams.Authorization, " ")
			fmt.Printf("\n数组长度：%d, token=>%s\n", len(token), token[1]) // 返回：2 ,token有效值

			context.Set("UserToken", token[1])
			fmt.Printf("中间件验证用户token：%s 完成！", token[1])
			context.Next()
		}

	}
}
