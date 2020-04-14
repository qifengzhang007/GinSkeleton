package Authorization

import (
	"fmt"
	"strings"

	//"GinSkeleton/Vendors/github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Token         string `header:"token"`
	Authorization string `header:"Authorization"`
}

func CheckAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//  模拟验证token
		//token:=context.DefaultQuery("token","")
		V_HeaderParams := HeaderParams{}
		//  Authorization  Bearer  在postman 等客户端软件传递的时候，都会被添加至Header头参数，键名 Authorization ，键值 Bearer Token  ....
		//  推荐使用ShouldbindHeader 方式获取头参数
		context.ShouldBindHeader(&V_HeaderParams)
		// Bearer:=context.Request.Header.Get("Authorization");  第二种方式从header头参数中获取键对应值

		token := strings.Split(V_HeaderParams.Authorization, " ")
		fmt.Printf("\n数组长度：%d, token=>%s\n", len(token), token[1]) // 返回：2 ,token有效值

		if len(V_HeaderParams.Authorization) < 1 {
			context.JSON(401,
				gin.H{
					"code": 401,
					"msg":  "NoAuthorazation",
				})
			//暂停执行
			context.Abort()
		} else {
			context.Set("UserToken", token[1])
			fmt.Printf("中间件验证用户token：%s 完成！", token[1])
			context.Next()
		}

	}
}
