package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Base
	Pass string `form:"pass" json:"pass"`
}

func (c *Login) CheckParams(context *gin.Context) {
	var v_form_params *Login = &Login{
		//&CodelistBase{},
	}
	if err := context.ShouldBind(v_form_params); err != nil {
		fmt.Printf("验证器出错")
		return
	}
	fmt.Printf("%#v\n", v_form_params)
	fmt.Println(v_form_params.Name)

	if len(v_form_params.Name) < 3 || len((*v_form_params).Pass) < 6 {
		fmt.Println("参数不符合规定，name、pass 长度有问题")
		return
	}
	// 验证完成，调用控制器
	(&Admin.Users{}).Login(context)
}

//  【必须的操作】 请记得去验证器工厂注册
