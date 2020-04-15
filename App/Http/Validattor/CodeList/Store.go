package CodeList

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Store struct {
	Stored_at time.Time `form:"Stored_at" binding:"required" time_format:"2006-01-02"`
	CodelistBase
}

func (c *Store) CheckParams(context *gin.Context) {
	var v_form_params = &Store{
		//CodelistBase:CodelistBase{},
	}
	context.ShouldBind(v_form_params)
	fmt.Println(v_form_params.Name, v_form_params.Code, v_form_params.Stored_at)
	if len(v_form_params.Name) > 1 && len(v_form_params.Code) == 6 {
		// 验证完成，调用控制器
		(&Admin.CodeList{}).Store(context)

	} else {

		return
	}
}
