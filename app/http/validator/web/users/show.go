package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	common_data_type "goskeleton/app/http/validator/common/data_type"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Show struct {
	// 表单参数验证结构体支持匿名结构体嵌套
	UserName string `form:"user_name" json:"user_name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	common_data_type.Page
}

// 验证器语法，参见 Register.go文件，有详细说明
func (s Show) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&s); err != nil {
		errs := gin.H{
			"tips": "UserShow参数校验失败，参数不符合规定，user_name（长度>0）、page的值(>0)、limits的值（>0)",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(s, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserShow表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Show(extraAddBindDataContext)
	}
}
