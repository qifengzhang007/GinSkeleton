package users

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/web"
	"goskeleton/app/utils/response"
	"strings"
)

type RefreshToken struct {
	Authorization string `json:"token" header:"Authorization" binding:"required,min=20"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (r RefreshToken) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBindHeader(&r); err != nil {
		// 将表单参数验证器出现的错误直接交给错误翻译器统一处理即可
		response.ValidatorError(context, err)
		return
	}
	token := strings.Split(r.Authorization, " ")
	if len(token) == 2 {
		context.Set(consts.ValidatorPrefix+"token", token[1])
		(&web.Users{}).RefreshToken(context)
	} else {
		errs := gin.H{
			"tips": "Token不合法，token请放置在header头部分，按照按=>键提交，例如：Authorization：Bearer 你的实际token....",
		}
		response.Fail(context, consts.JwtTokenFormatErrCode, consts.JwtTokenFormatErrMsg, errs)
	}

}
