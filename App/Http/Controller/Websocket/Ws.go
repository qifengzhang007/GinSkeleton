package Websocket

import (
	service_ws "GinSkeleton/App/Service/Websocket"
	"github.com/gin-gonic/gin"
)

/**
websocket 想要了解更多具体细节请参见以下文档
文档地址：https://github.com/gorilla/websocket/tree/master/examples
*/

type Ws struct {
}

// OnOpen 主要解决握手+协议升级
func (w *Ws) OnOpen(context *gin.Context) (*service_ws.Ws, bool) {
	return (&service_ws.Ws{}).OnOpen(context)
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(service_ws *service_ws.Ws, context *gin.Context) {
	service_ws.OnMessage(context)
}
