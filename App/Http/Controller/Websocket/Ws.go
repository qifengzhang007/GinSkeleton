package Websocket

import (
	ServiceWebsocket "GinSkeleton/App/Service/Websocket"
	"github.com/gin-gonic/gin"
)

/**
2020-04-28  本段代码开发中，未完成，请勿使用，预计2020-05-01放假之前完成
*/

type Ws struct {
}

// 该函数基本上不需要做任何业务，都是在做系统初始化websocet连接的一些东西
func (w *Ws) OnOpen(context *gin.Context) bool {

	return (&ServiceWebsocket.Ws{}).OnOpen(context)

}

func (w *Ws) OnMessage(context *gin.Context) {

}

func (w *Ws) OnClose() {

}
