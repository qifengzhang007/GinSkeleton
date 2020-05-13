package Home

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Api"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 门户类前端接口模拟一个获取新闻的参数验证器
type News struct {
	NewsType string  `form:"newstype" json:"newstype"  binding:"required,min=1"` //  验证规则：必填，最小长度为1
	Page     float64 `form:"page" json:"page"  binding:"required,min=1"`         //  验证规则：必填，最小值为1（float类型，min=1代表最小值为1）
	Limit    float64 `form:"limit" json:"limit"  binding:"required,min=1"`       //  验证规则：必填，最小值为1（float类型，min=1代表最小值为1）
}

func (n News) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&n); err != nil {
		errs := gin.H{
			"tips": "HomeNews参数校验失败，参数不符合规定，limit(长度>=1)、page>=1、limit>=1,请按照规则自己检查",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(n, Consts.Validator_Prefix, context)
	if extraAddBindDataContext == nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",HomeNews表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&Api.Home{}).News(extraAddBindDataContext)
	}

}
