package CodeList

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Delete struct {
	Id int `form:"id" binding:"required"`
}

func (c *Delete) CheckParams(context *gin.Context) {

	var v_form_delete = &Delete{}
	context.ShouldBind(v_form_delete)

	fmt.Println(v_form_delete.Id)

	if v_form_delete.Id > 0 {
		// 验证完成，调用控制器
		(&Admin.CodeList{}).Delete(context)

	} else {

		return
	}
}
