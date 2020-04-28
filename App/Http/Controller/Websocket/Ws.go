package Websocket

import (
	service_ws "GinSkeleton/App/Service/Websocket"
	"github.com/gin-gonic/gin"
)

type Ws struct {
}

// OnOpne 主要解决握手+协议升级
func (w *Ws) OnOpen(context *gin.Context) (*service_ws.Ws, bool) {
	return (&service_ws.Ws{}).OnOpen(context)
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(service_ws *service_ws.Ws, context *gin.Context) {
	service_ws.OnMessage(context)
}
