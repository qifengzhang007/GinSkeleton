package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Web"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 验证器是本项目骨架的先锋队，必须发挥它的极致优势，具体参考地址：
//https://godoc.org/github.com/go-playground/validator   ,该验证器非常强大，强烈建议重点发挥，
//请求正式进入控制器等后面的业务逻辑层之前，参数的校验必须在验证器层完成，后面的控制器等就只管获取各种参数，代码一把梭

// 给出一些最常用的验证规则：
//required  必填；
//len=11 长度=11；
//min=3  如果是数字，验证的是数据范围，最小值为3，如果是文本，验证的是最小长度为3，
//max=6 如果是数字，验证的是数字最大值为6，如果是文本，验证的是最大长度为6
// mail 验证邮箱
//gt=3  对于文本就是长度>=3
//lt=6  对于文本就是长度<=6

type Register struct {
	Base
	Phone string `form:"phone" json:"phone"  binding:"required,len=11"`    //  验证规则：必填，长度必须=11
	Pass  string `form:"pass" json:"pass" binding:"required,min=3,max=20"` //必填，密码长度范围：【3,20】闭区间
	//CardNo  string `form:"card_no" json:"card_no" binding:"required,len=18"`	//身份证号码，必填，长度=18
}

func (r Register) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&r); err != nil {
		errs := gin.H{
			"tips": "UserRegister参数校验失败，参数不符合规定，name长度(>=1)、pass长度(3,6)、Phone（=11位），不允许注册",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}
	//2.继续验证具有中国特色的参数，例如 身份证号码等，基本语法校验了长度18位，然后可以自行编写正则表达式等更进一步验证每一部分组成
	// r.CardNo  获取值继续校验，这里省略.....

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(r, Consts.Validator_Prefix, context)
	if extraAddBindDataContext == nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&Web.Users{}).Register(extraAddBindDataContext)
	}

}
