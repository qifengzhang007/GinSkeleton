package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/dataTransfer"
	"goskeleton/app/utils/response"
	"net/http"
)

type Show struct {
	Base
	Page   float64 `form:"page" json:"page" binding:"required,gt=0"`     // 必填，页面值>0
	Limits float64 `form:"limits" json:"limits" binding:"required,gt=0"` // 必填，每页条数值>0
}

// 验证器语法，参见 Register.go文件，有详细说明
func (s Show) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&s); err != nil {
		errs := gin.H{
			"tips": "UserShow参数校验失败，参数不符合规定，name（不能为空）、page的值(>0)、limits的值（>0)",
			"err":  err.Error(),
		}
		response.ReturnJson(context, http.StatusBadRequest, consts.Validator_ParamsCheck_Fail_Code, consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := dataTransfer.DataAddContext(s, consts.Validator_Prefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnJson(context, http.StatusInternalServerError, consts.Server_Occurred_Error_Code, consts.Server_Occurred_Error_Msg+",UserShow表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Show(extraAddBindDataContext)
	}
}
