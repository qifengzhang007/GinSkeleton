package CodeList

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Show struct {
	CodelistBase
}

func (c *Show) CheckParams(context *gin.Context) {
	var v_form_params *Show = &Show{
		//&CodelistBase{},
	}
	if err := context.ShouldBind(v_form_params); err != nil {
		fmt.Printf("验证器出错")
		return
	}
	if len(v_form_params.Name) < 3 || len(v_form_params.Code) != 6 {
		fmt.Println("参数不符合规定，name、code长度有问题")
		return
	}
	// 验证完成，调用控制器
	(&Admin.CodeList{}).Show(context)
}
