package Websocket

import (
	service_ws "GinSkeleton/App/Service/Websocket"
	"github.com/gin-gonic/gin"
)

type Ws struct {
}

// OnOpne 主要解决握手+协议升级
func (w *Ws) OnOpen(context *gin.Context) (*service_ws.Ws, bool) {
	serv_ws := &service_ws.Ws{}
	return serv_ws, serv_ws.OnOpen(context)
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(serv_ws *service_ws.Ws, context *gin.Context) {
	serv_ws.OnMessage(context)
}
