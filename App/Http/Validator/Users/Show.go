package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Show struct {
	Base
	Page   float64 `form:"page" json:"page" binding:"required"`
	Limits float64 `form:"limits" json:"limits" binding:"required"`
}

func (s *Show) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&s); err != nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserShow验证器绑定出错", "")
		return
	}

	if len(s.Name) > 0 && s.Page > 0 && s.Limits > 0 {
		//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
		extraAddBindDataContext := DaTaTransfer.DataAddContext(s, Consts.Validator_Prefix, context)
		if extraAddBindDataContext == nil {
			Response.ReturnJson(context, http.StatusBadRequest, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserShow表单验证器json化失败", "")
		} else {
			// 验证完成，调用控制器,并将追加标案参数验证器的上下文传递给控制器
			(&Admin.Users{}).Show(extraAddBindDataContext)
		}
	} else {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg+"，UserShow参数校验失败，参数不符合规定，name（不能为空）、page(>0)、limits（>0)", "")
	}

}
