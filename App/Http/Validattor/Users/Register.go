package Users

import (
	"GinSkeleton/App/Global/Consts"
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

func (r *Register) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&r); err != nil {
		fmt.Printf("验证器出错")
		return
	}

	if len(r.Name) < 3 || len(r.Pass) < 6 || len(r.Phone) != 11 {
		fmt.Println("参数不符合规定，name、pass、Phone 长度有问题，不允许注册")
		return
	}
	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(r, Consts.Validattor_Prefix, context)
	if extraAddBindDataContext == nil {
		fmt.Println("表单参数验证器json化失败..")
	} else {
		(&Admin.Users{}).Register(extraAddBindDataContext)
	}
}

//  请记得将表单验证器注册在容器工厂
