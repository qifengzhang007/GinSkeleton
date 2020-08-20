package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
	"net/http"
)

type Login struct {
	Base
	Pass string `form:"pass" json:"pass" binding:"required,min=6,max=20"` //  密码为 必填，长度>=6
}

// 验证器语法，参见 Register.go文件，有详细说明

func (l Login) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&l); err != nil {
		errs := gin.H{
			"tips": "UserRegister参数校验失败，参数不符合规定，name、pass、Phone 长度有问题，不允许登录",
			"err":  err.Error(),
		}
		response.ReturnJson(context, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(l, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnJson(context, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+",userLogin表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Login(extraAddBindDataContext)
	}

}
