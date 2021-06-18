package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Update struct {
	BaseField
	Id
	// 表单参数验证结构体支持匿名结构体嵌套、以及匿名结构体与普通字段组合
	RealName string `form:"real_name" json:"real_name" binding:"required,min=2"`
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	Remark   string `form:"remark" json:"remark"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (u Update) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&u); err != nil {
		errs := gin.H{
			"tips": "UserUpdate，参数校验失败，请检查id(>0),user_name(>=1)、pass(>=6)、real_name(>=2)、phone长度(=11)",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(u, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserUpdate表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Update(extraAddBindDataContext)
	}
}
