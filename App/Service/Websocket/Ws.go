package Websocket

import (
	"GinSkeleton/App/Utils/Websocket/Core"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 1.将http连接升级到websocket连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024 * 2,
	WriteBufferSize: 65535,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Ws struct {
}

func (w *Ws) OnOpen(context *gin.Context) bool {

	// 2.将http协议升级到websocket协议.初始化一个有效的websocket长连接客户端
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Println(err)
		return false
	} else {
		// 初始化完成一个ws连接
		client := &Core.Client{
			Hub:  &Core.Hub{},
			Conn: conn,
			Send: make(chan []byte, 256),
		}
		client.Hub.Register <- client
		go client.ReadPump()
		go client.WritePump()
		return true
	}

}

func (w *Ws) OnMessage(context *gin.Context) {

}

func (w *Ws) OnClose() {

}
