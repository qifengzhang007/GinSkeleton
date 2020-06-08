package Users

import (
	"GinSkeleton/App/Global/Consts"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 模拟Aop 实现对某个控制器函数的前置和后置回调

type DestroyAfter struct{}

func (d *DestroyAfter) After(context *gin.Context) {
	// 后置函数可以使用异步执行
	go func() {
		userId := context.GetFloat64(Consts.Validator_Prefix + "id")
		fmt.Printf("模拟 Users 删除操作， After 回调,用户ID：%.f\n", userId)
	}()
}
