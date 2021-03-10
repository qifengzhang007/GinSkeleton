package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goskeleton/app/http/validator/web/users/data_type"
	"goskeleton/app/utils/response"
)

type Test struct {
	data_type.StuInfo
}

func (l Test) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBindJSON(&l); err != nil {
		errs := gin.H{
			"tips": "UserRegister参数校验失败，参数不符合规定，user_name、pass、 长度有问题，不允许登录",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}
	fmt.Printf("%#+v\n", l.StuInfo)
	fmt.Printf("%#+v\n", l)
	context.Set("StuInfo", l.StuInfo)

	if val, ok := context.Get("StuInfo"); ok {
		if stu, ok := val.(data_type.StuInfo); ok {
			fmt.Printf("数据断言OK,获取值：%#+v\n", stu)
		}
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	//extraAddBindDataContext := data_transfer.DataAddContext(l, consts.ValidatorPrefix, context)
	//if extraAddBindDataContext == nil {
	//	response.ErrorSystem(context, "userTest表单验证器json化失败", "")
	//} else {
	//	// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
	//	//(&web.Users{}).Test(extraAddBindDataContext)
	//}

}
