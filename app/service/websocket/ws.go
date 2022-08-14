package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/websocket/core"
)

/**
websocket模块相关事件执行顺序：
1.onOpen
2.OnMessage
3.OnError
4.OnClose
*/

type Ws struct {
	WsClient *core.Client
}

// OnOpen 事件函数
func (w *Ws) OnOpen(context *gin.Context) (*Ws, bool) {
	if client, ok := (&core.Client{}).OnOpen(context); ok {

		token := context.GetString(consts.ValidatorPrefix + "token")
		variable.ZapLog.Info("获取到的客户端上线时携带的唯一标记值：", zap.String("token", token))

		// 成功上线以后，开发者可以基于客户端上线时携带的唯一参数(这里用token键表示)
		// 在数据库查询更多的其他字段信息，直接追加在 Client 结构体上，方便后续使用
		//client.ClientMoreParams.UserParams1 = "123"
		//client.ClientMoreParams.UserParams2 = "456"
		//fmt.Printf("最终每一个客户端(client) 已有的参数：%+v\n", client)

		w.WsClient = client
		go w.WsClient.Heartbeat() // 一旦握手+协议升级成功，就为每一个连接开启一个自动化的隐式心跳检测包
		return w, true
	} else {
		return nil, false
	}
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(context *gin.Context) {
	go w.WsClient.ReadPump(func(messageType int, receivedData []byte) {
		//参数说明
		//messageType 消息类型，1=文本
		//receivedData 服务器接收到客户端（例如js客户端）发来的的数据，[]byte 格式

		tempMsg := "服务器已经收到了你的消息==>" + string(receivedData)
		// 回复客户端已经收到消息;
		if err := w.WsClient.SendMessage(messageType, tempMsg); err != nil {
			variable.ZapLog.Error("消息发送出现错误", zap.Error(err))
		}

	}, w.OnError, w.OnClose)
}

// OnError 客户端与服务端在消息交互过程中发生错误回调函数
func (w *Ws) OnError(err error) {
	w.WsClient.State = 0 // 发生错误，状态设置为0, 心跳检测协程则自动退出
	variable.ZapLog.Error("远端掉线、卡死、刷新浏览器等会触发该错误:", zap.Error(err))
	//fmt.Printf("远端掉线、卡死、刷新浏览器等会触发该错误: %v\n", err.Error())
}

// OnClose 客户端关闭回调，发生onError回调以后会继续回调该函数
func (w *Ws) OnClose() {

	w.WsClient.Hub.UnRegister <- w.WsClient // 向hub管道投递一条注销消息，由hub中心负责关闭连接、删除在线数据
}

// GetOnlineClients  获取在线的全部客户端
func (w *Ws) GetOnlineClients() {

	fmt.Printf("在线客户端数量：%d\n", len(w.WsClient.Hub.Clients))
}

// BroadcastMsg  (每一个客户端都有能力)向全部在线客户端广播消息
func (w *Ws) BroadcastMsg(sendMsg string) {
	for onlineClient := range w.WsClient.Hub.Clients {

		//获取每一个在线的客户端，向远端发送消息
		if err := onlineClient.SendMessage(websocket.TextMessage, sendMsg); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsWebsocketWriteMgsFail, zap.Error(err))
		}
	}
}
