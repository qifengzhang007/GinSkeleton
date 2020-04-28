package Websocket

import (
	"GinSkeleton/App/Global/Consts"
	controller_ws "GinSkeleton/App/Http/Controller/Websocket"
	"GinSkeleton/App/Utils/Config"
	"GinSkeleton/App/Utils/Response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Connect struct {
	Token string `form:"token" binding:"required,min=10"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (h *Connect) CheckParams(context *gin.Context) {

	// 1. 首先检查是否开启websocket服务配置（在配置项中开启）
	if Config.CreateYamlFactory().GetInt("Websocket.Start") != 1 {
		fmt.Println("111111")
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Ws_Server_Not_Start_Code, Consts.Ws_Server_Not_Start_Msg, "")
		return
	}
	//2.基本的验证规则没有通过
	if err := context.ShouldBind(h); err != nil {
		errs := gin.H{
			"tips": "请在get参数中提交token信息",
			"err":  err.Error(),
		}
		fmt.Println(err.Error())
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	if service_ws, ok := (&controller_ws.Ws{}).OnOpen(context); ok == false {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Ws_Open_Fail_Code, Consts.Ws_Open_Fail_Msg, "")
	} else {
		(&controller_ws.Ws{}).OnMessage(service_ws, context)
	}
}
