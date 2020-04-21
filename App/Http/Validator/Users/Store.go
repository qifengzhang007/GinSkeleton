package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Store struct {
	Base
	//Created_at time.Time `form:"Stored_at" binding:"required" time_format:"2006-01-02"`
	Pass      string `form:"pass" json:"pass" binding:"required"`
	Real_name string `form:"real_name" json:"real_name" binding:"required"`
	Phone     string `form:"phone" json:"phone" binding:"required"`
	Remark    string `form:"remark" json:"remark" binding:"required"`
}

func (s *Store) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(&s); err != nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserStore验证器绑定出错", "")
		return
	}

	if len(s.Name) >= 6 && len(s.Pass) >= 6 && len(s.Phone) == 11 && len(s.Real_name) > 2 {
		//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
		extraAddBindDataContext := DaTaTransfer.DataAddContext(s, Consts.Validator_Prefix, context)
		if extraAddBindDataContext == nil {
			Response.ReturnJson(context, http.StatusBadRequest, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserStore表单验证器json化失败", "")
		} else {
			// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
			(&Admin.Users{}).Store(extraAddBindDataContext)
		}

	} else {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg+"，UserShow参数校验失败，参数校验失败，请检查账号(>=6)、密码(>=6)、姓名(>=2)、手机号码(=11)长度", "")
	}
}
