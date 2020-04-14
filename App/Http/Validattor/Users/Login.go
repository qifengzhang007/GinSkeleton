package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UsersLogin struct {
	Base
	Pass string `form:"pass" json:"pass"`
}

func (c *UsersLogin) CheckParams(context *gin.Context) {
	var v_form_params *UsersLogin = &UsersLogin{
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
