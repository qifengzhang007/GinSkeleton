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

// onOpen 基本不需要做什么
func (w *Ws) OnOpen(context *gin.Context) (*Ws, bool) {
	if client, ok := (&Core.Client{}).OnOpen(context); ok {
		w.WsClient = client
		return w, true
	} else {
		return nil, false
	}
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(context *gin.Context) {
	go w.WsClient.ReadPump(func(messageType int, received_data []byte) {
		//参数说明
		//messageType 消息类型，1=文本
		//received_data 服务器接收到客户端（例如js客户端）发来的的数据，[]byte 格式

		w.WsClient.Conn.WriteMessage(websocket.TextMessage, append([]byte("hello，服务器已经收到你的消息=>"), []byte(received_data)...)) // 回复客户端已经收到消息

	}, w.OnError, w.OnClose)
	go w.WsClient.Heartbeat(w.OnClose)
}

// OnError 客户端与服务端在消息交互过程中发生错误回调函数
func (w *Ws) OnError(err error) {

	//err 接收消息期间发生的错误（一般来说，远端掉线、断开、卡死等会触发此错误）
	fmt.Printf("ws连接消息轮训期间发生错误: %v", err.Error())
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		log.Printf("error: %v", err)
	}
}

// OnClose 客户端关闭回调（主要有：下线、断开、卡死等无法正常通讯的状况）
func (w *Ws) OnClose() {
	w.WsClient.Hub.UnRegister <- w.WsClient // 向hub管道投递一条注销消息，有hub中心负责关闭连接、删除在线数据
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
