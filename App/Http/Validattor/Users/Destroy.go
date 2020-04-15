package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Destroy struct {
	Id int `form:"id" binding:"required"`
}

func (c *Destroy) CheckParams(context *gin.Context) {

	var v_form_Destroy = &Destroy{}
	context.ShouldBind(v_form_Destroy)

	fmt.Println(v_form_Destroy.Id)

	if v_form_Destroy.Id > 0 {
		// 验证完成，调用控制器
		(&Admin.Users{}).Destroy(context)

	} else {

		return
	}
}
