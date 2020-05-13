package Websocket

import (
	"GinSkeleton/App/Global/Consts"
	controller_ws "GinSkeleton/App/Http/Controller/Websocket"
	"GinSkeleton/App/Utils/Config"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Connect struct {
	Token string `form:"token" binding:"required,min=10"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (c Connect) CheckParams(context *gin.Context) {

	// 1. 首先检查是否开启websocket服务配置（在配置项中开启）
	if Config.CreateYamlFactory().GetInt("Websocket.Start") != 1 {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Ws_Server_Not_Start_Code, Consts.Ws_Server_Not_Start_Msg, "")
		return
	}
	//2.基本的验证规则没有通过
	if err := context.ShouldBind(&c); err != nil {
		errs := gin.H{
			"tips": "请在get参数中提交token信息,demo格式：ws://127.0.0.1:2020?token=asasasaasasasssddsdsd",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	if service_ws, ok := (&controller_ws.Ws{}).OnOpen(context); ok == false {
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Ws_Open_Fail_Code, Consts.Ws_Open_Fail_Msg, "")
	} else {
		(&controller_ws.Ws{}).OnMessage(service_ws, context) // 注意这里传递的service_ws必须是调用open返回的，必须保证的ws对象的一致性
	}
}
