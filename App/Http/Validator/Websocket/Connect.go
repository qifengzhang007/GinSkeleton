package Websocket

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Websocket"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Connect struct {
	Authorization string `header:"Authorization" binding:"required,min=10"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (h *Connect) CheckParams(context *gin.Context) {

	//1.基本的验证规则没有通过
	if err := context.ShouldBindHeader(h); err != nil {
		errs := gin.H{
			"tips": "请在header头添加Authorization:Beaer (你的Token.....)进行简单的验证",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	if (&Websocket.Ws{}).OnOpen(context) == false {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Ws_Open_Fail_Code, Consts.Ws_Open_Fail_Msg, "")
	} else {
		(&Websocket.Ws{}).OnMessage(context)
	}
}
