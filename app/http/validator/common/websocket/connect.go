package websocket

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	controllerWs "goskeleton/app/http/controller/websocket"
	"goskeleton/app/utils/response"
	"goskeleton/app/utils/yml_config"
	"net/http"
)

type Connect struct {
	Token string `form:"token" binding:"required,min=10"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (c Connect) CheckParams(context *gin.Context) {

	// 1. 首先检查是否开启websocket服务配置（在配置项中开启）
	if yml_config.CreateYamlFactory().GetInt("Websocket.Start") != 1 {
		response.ReturnJson(context, http.StatusBadRequest, consts.WsServerNotStartCode, consts.WsServerNotStartMsg, "")
		return
	}
	//2.基本的验证规则没有通过
	if err := context.ShouldBind(&c); err != nil {
		errs := gin.H{
			"tips": "请在get参数中提交token信息,demo格式：ws://127.0.0.1:2020?token=this is a series token...",
			"err":  err.Error(),
		}
		response.ReturnJson(context, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, errs)
		return
	}

	if serviceWs, ok := (&controllerWs.Ws{}).OnOpen(context); ok == false {
		response.ReturnJson(context, http.StatusBadRequest, consts.WsOpenFailCode, consts.WsOpenFailMsg, "")
	} else {
		(&controllerWs.Ws{}).OnMessage(serviceWs, context) // 注意这里传递的service_ws必须是调用open返回的，必须保证的ws对象的一致性
	}
}
