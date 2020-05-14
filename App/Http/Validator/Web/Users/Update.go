package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Web"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Update struct {
	Base
	Id        float64 `form:"id" json:"id" binding:"required,min=1"` // 注意：gin框架数字的存储形式都是 float64
	Pass      string  `form:"pass" json:"pass" binding:"required,min=6"`
	Real_name string  `form:"real_name" json:"real_name" binding:"required,min=2"`
	Phone     string  `form:"phone" json:"phone" binding:"required,len=11"`
	Remark    string  `form:"remark" json:"remark"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (u Update) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&u); err != nil {
		errs := gin.H{
			"tips": "UserUpdate，参数校验失败，请检查id(>0),name(>=1)、pass(>=6)、real_name(>=2)、phone长度(=11)",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(u, Consts.Validator_Prefix, context)
	if extraAddBindDataContext == nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserUpdate表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&Web.Users{}).Update(extraAddBindDataContext)
	}
}
