package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Web"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RefreshToken struct {
	Authorization string `json:"token" header:"Authorization" binding:"required,min=20"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (r RefreshToken) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBindHeader(&r); err != nil {
		errs := gin.H{
			"tips": "Token参数校验失败，参数不符合规定，token 长度>=20",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}
	token := strings.Split(r.Authorization, " ")
	if len(token) == 2 {
		context.Set(Consts.Validator_Prefix+"token", token[1])
		(&Web.Users{}).RefreshToken(context)
	} else {
		errs := gin.H{
			"tips": "Token不合法，token请放置在header头部分，按照按=>键提交，例如：Authorization：Bearer 你的实际token....",
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
	}

}
