package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
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
	BaseField
	// 表单参数验证结构体支持匿名结构体嵌套、以及匿名结构体与普通字段组合
	Phone  string `form:"phone" json:"phone"`     // 手机号， 非必填
	CardNo string `form:"card_no" json:"card_no"` //身份证号码，非必填
}

// 特别注意: 表单参数验证器结构体的函数，绝对不能绑定在指针上
// 我们这部分代码项目启动后会加载到容器，如果绑定在指针，一次请求之后，会造成容器中的代码段被污染

func (r Register) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&r); err != nil {
		errs := gin.H{
			"tips": "UserRegister参数校验失败，参数不符合规定，user_name 长度(>=1)、pass长度[6,20]、不允许注册",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}
	//2.继续验证具有中国特色的参数，例如 身份证号码等，基本语法校验了长度18位，然后可以自行编写正则表达式等更进一步验证每一部分组成
	// r.CardNo  获取值继续校验，这里省略.....

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(r, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Register(extraAddBindDataContext)
	}

}
