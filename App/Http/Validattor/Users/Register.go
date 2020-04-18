package Users

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validattor/Core/DaTaTransfer"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Base
	Phone string `form:"phone" json:"phone"  bind:"required"`
	Pass  string `form:"pass" json:"pass" bind:"required"`
}

func (c *Register) CheckParams(context *gin.Context) {
	var v_form_params *Register = &Register{}
	if err := context.ShouldBind(v_form_params); err != nil {
		fmt.Printf("验证器出错")
		return
	}
	fmt.Printf("%#v\n", v_form_params)
	fmt.Println(v_form_params.Name)

	if len(v_form_params.Name) < 3 || len((*v_form_params).Pass) < 6 || len((*v_form_params).Phone) != 11 {
		fmt.Println("参数不符合规定，name、pass、Phone 长度有问题，不允许注册")
		return
	}
	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindData := DaTaTransfer.DataAddContext(v_form_params, context)
	(&Admin.Users{}).Register(extraAddBindData)
}

//  请记得将表单验证器注册在容器工厂
