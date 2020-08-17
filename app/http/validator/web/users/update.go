package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
	"net/http"
)

type Update struct {
	Base
	Id       float64 `form:"id" json:"id" binding:"required,min=1"` // 注意：gin框架数字的存储形式都是 float64
	Pass     string  `form:"pass" json:"pass" binding:"required,min=6"`
	RealName string  `form:"real_name" json:"real_name" binding:"required,min=2"`
	Phone    string  `form:"phone" json:"phone" binding:"required,len=11"`
	Remark   string  `form:"remark" json:"remark"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (u Update) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&u); err != nil {
		errs := gin.H{
			"tips": "UserUpdate，参数校验失败，请检查id(>0),name(>=1)、pass(>=6)、real_name(>=2)、phone长度(=11)",
			"err":  err.Error(),
		}
		response.ReturnJson(context, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(u, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnJson(context, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+",UserUpdate表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Update(extraAddBindDataContext)
	}
}
