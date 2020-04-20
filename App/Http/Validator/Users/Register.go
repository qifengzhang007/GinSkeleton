package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Register struct {
	Base
	Phone string `form:"phone" json:"phone"  bind:"required"`
	Pass  string `form:"pass" json:"pass" bind:"required"`
}

func (r *Register) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&r); err != nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserRegister验证器绑定出错", "")
		return
	}

	if len(r.Name) >= 3 && len(r.Pass) >= 6 && len(r.Phone) == 11 {
		//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
		extraAddBindDataContext := DaTaTransfer.DataAddContext(r, Consts.Validator_Prefix, context)
		if extraAddBindDataContext == nil {
			Response.ReturnJson(context, http.StatusBadRequest, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserRegister表单验证器json化失败", "")
		} else {
			(&Admin.Users{}).Register(extraAddBindDataContext)
		}
	} else {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg+"，UserRegister参数校验失败，参数不符合规定，name、pass、Phone 长度有问题，不允许注册", "")
	}

}
