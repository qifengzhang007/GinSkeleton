package users

import (
	"ginskeleton/app/global/consts"
	"ginskeleton/app/global/variable"
	"github.com/gin-gonic/gin"
)

// 模拟Aop 实现对某个控制器函数的前置和后置回调

type DestroyAfter struct{}

func (d *DestroyAfter) After(context *gin.Context) {
	// 后置函数可以使用异步执行
	go func() {
		userId := context.GetFloat64(consts.ValidatorPrefix + "id")
		variable.ZapLog.Sugar().Infof("模拟 Users 删除操作， After 回调,用户ID：%.f\n", userId)
	}()
}
