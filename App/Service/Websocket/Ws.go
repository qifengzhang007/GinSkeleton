package Websocket

import (
	"GinSkeleton/App/Utils/Websocket/Core"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type Ws struct {
	WsClient *Core.Client
}

func (w *Ws) OnOpen(context *gin.Context) bool {
	if client, ok := (&Core.Client{}).OnOpen(context); ok {
		w.WsClient = client
		return true
	} else {
		return false
	}
}

func (w *Ws) OnMessage(context *gin.Context) {

	go w.WsClient.ReadPump(func(messageType int, p []byte) {
		//参数说明
		//messageType 接收到的消息类型
		//p 接收到的消息，[]byte 格式

		w.WsClient.Send <- p // 服务器收到的原始消息,首先放入管道，起到一个缓冲作用
		w.WsClient.Conn.WriteMessage(websocket.TextMessage, append([]byte("hello，服务器已经收到你的消息===>"), []byte(p)...))
		n := len(w.WsClient.Send)
		// 防止管道存在堆积的消息，用for取出消息，逐条发送给远端，
		for i := 0; i < n; i++ {
			//向远端发送消息
			w.WsClient.Conn.WriteMessage(websocket.TextMessage, append([]byte("hello，服务器已经收到你的消息===>"), <-w.WsClient.Send...))
		}
		//		close(w.WsClient.Send)

	}, w.OnError, w.OnClose)
}

func (w *Ws) OnError(err error) {

	//err 接收消息期间发生的错误（一般来说，远端掉线、断开、卡死等会触发此错误）
	fmt.Printf("ws连接消息轮训期间发生错误: %v", err.Error())
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		log.Printf("error: %v", err)
	}
}

func (w *Ws) OnClose() {
	w.WsClient.Hub.UnRegister <- w.WsClient // 向hub注销一个下线的ws连接
	w.WsClient.Conn.Close()

}

//获取在线的全部客户端
func (w *Ws) GetOnlineClients() {
	fmt.Printf("在线客户端数量：%d\n", len(w.WsClient.Hub.Clients))
}

// 向全部在线客户端广播消息
func (w *Ws) BroadcastMsg(send_msg string) {
	//获取每一个在线的客户端，向远端发送消息
	for online_client, _ := range w.WsClient.Hub.Clients {
		online_client.Conn.WriteMessage(websocket.TextMessage, []byte(send_msg))
	}
}
