package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ShowList struct {
	Base
}

func (c *ShowList) CheckParams(context *gin.Context) {

	fmt.Printf("验证器获取的参数:\n%#v\n", context.Query("username"))
	if len(context.Query("username")) >= 1 {
		// 验证完成，调用控制器
		(&Admin.Users{}).ShowList(context)

	} else {
		context.JSON(-100, gin.H{
			"code": -100,
			"msg":  "参数校验失败，请检查 username 关键词是否有效",
		})
		fmt.Printf("这里不会执行！！！")
	}
}
