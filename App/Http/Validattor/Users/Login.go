package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validattor/Core/DaTaTransfer"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Base
	Pass string `form:"pass" json:"pass"`
}

func (l *Login) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&l); err != nil {
		fmt.Printf("验证器出错")
		return
	}
	if len(l.Name) < 3 || len(l.Pass) < 6 {
		fmt.Println("参数长度不符合规定，name长度>=3,pass长度>=6")
		return
	}
	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(l, Consts.Validattor_Prefix, context)
	if extraAddBindDataContext == nil {
		fmt.Println("表单参数验证器json化失败..")
	} else {
		// 验证完成，调用控制器,并将追加标案参数验证器的上下文传递给控制器
		(&Admin.Users{}).Login(extraAddBindDataContext)
	}
}

//  【必须的操作】 请记得去验证器工厂注册
