package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Base
	Pass string `form:"pass" json:"pass"`
}

func (l *Login) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&l); err != nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",Userlogin验证器绑定出错", "")
		return
	}

	if len(l.Name) >= 3 && len(l.Pass) >= 6 {
		//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
		extraAddBindDataContext := DaTaTransfer.DataAddContext(l, Consts.Validator_Prefix, context)
		if extraAddBindDataContext == nil {
			Response.ReturnJson(context, http.StatusBadRequest, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",Userlogin表单验证器json化失败", "")
		} else {
			// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
			(&Admin.Users{}).Login(extraAddBindDataContext)
		}
	} else {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg+"，Userlogin参数校验失败，参数长度不符合规定，name长度>=3,pass长度>=6", "")
	}

}
