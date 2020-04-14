package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Create struct {
	Base
	Created_at time.Time `form:"created_at" binding:"required" time_format:"2006-01-02"`
	Pass       string    `form:"pass" json:"pass" binding:"required"`
	Remark     string    `form:"remark" json:"remark" binding:"required"`
}

func (c *Create) CheckParams(context *gin.Context) {
	var v_form_create = &Create{
		//CodelistBase:CodelistBase{},
	}
	context.ShouldBind(v_form_create)
	fmt.Printf("验证器获取的参数:\n%#v\n", v_form_create)
	if len(v_form_create.Name) >= 6 && len(v_form_create.Pass) >= 6 {
		// 验证完成，调用控制器
		(&Admin.Users{}).Create(context)

	} else {
		context.JSON(-100, gin.H{
			"code": -100,
			"msg":  "参数校验失败，请检查账号、密码长度",
		})
		fmt.Printf("这里不会执行！！！")
	}
}
