package websocket

import (
	"github.com/gin-gonic/gin"
	serviceWs "goskeleton/app/service/websocket"
)

/**
websocket 想要了解更多具体细节请参见以下文档
文档地址：https://github.com/gorilla/websocket/tree/master/examples
*/

type Ws struct {
}

// OnOpen 主要解决握手+协议升级
func (w *Ws) OnOpen(context *gin.Context) (*serviceWs.Ws, bool) {
	return (&serviceWs.Ws{}).OnOpen(context)
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(serviceWs *serviceWs.Ws, context *gin.Context) {
	serviceWs.OnMessage(context)
}
