package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Web"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Store struct {
	Base
	//Created_at time.Time `form:"Stored_at" binding:"required" time_format:"2006-01-02"`
	Pass      string `form:"pass" json:"pass" binding:"required,min=6"`
	Real_name string `form:"real_name" json:"real_name" binding:"required,min=2"`
	Phone     string `form:"phone" json:"phone" binding:"required,len=11"`
	Remark    string `form:"remark" json:"remark" `
}

// 验证器语法，参见 Register.go文件，有详细说明

func (s Store) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&s); err != nil {
		errs := gin.H{
			"tips": "UserStore参数校验失败，参数校验失败，请检查name(>=1)、pass(>=6)、real_name(>=2)、phone(=11)",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(s, Consts.Validator_Prefix, context)
	if extraAddBindDataContext == nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserStore表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&Web.Users{}).Store(extraAddBindDataContext)
	}
}
