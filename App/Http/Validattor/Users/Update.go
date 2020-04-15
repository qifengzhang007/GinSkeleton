package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Update struct {
	Base
	Id         int       `form:"id" binding:"required"`
	Updated_at time.Time `form:"updated_at" binding:"required" time_format:"2006-01-02"`
}

func (c *Update) CheckParams(context *gin.Context) {

	var v_form_params = &Update{}
	context.ShouldBind(v_form_params)

	fmt.Println(v_form_params.Name, v_form_params.Name, (v_form_params.Updated_at).Format("2006-01-02"), v_form_params.Id)

	if len(v_form_params.Name) > 1 && len(v_form_params.Name) == 6 {
		// 验证完成，调用控制器
		context.Set("test_key", "测试在中间件添加额外键值")
		(&Admin.Users{}).Update(context)

	} else {

		return
	}
}
