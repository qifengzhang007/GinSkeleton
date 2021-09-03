###    websocket  

##### 1.基本用法  
> 以下代码展示的是每一个 websocket 客户端连接到服务端所拥有的功能
- [相关代码位置](../app/service/websocket/ws.go)  
```code 
package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
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

// onOpen 基本不需要做什么
func (w *Ws) OnOpen(context *gin.Context) (*Ws, bool) {
	if client, ok := (&core.Client{}).OnOpen(context); ok {
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

//获取在线的全部客户端
func (w *Ws) GetOnlineClients() {

	fmt.Printf("在线客户端数量：%d\n", len(w.WsClient.Hub.Clients))
}

// (每一个客户端都有能力)向全部在线客户端广播消息
func (w *Ws) BroadcastMsg(sendMsg string) {
	for onlineClient := range w.WsClient.Hub.Clients {

		//获取每一个在线的客户端，向远端发送消息
		if err := onlineClient.SendMessage(websocket.TextMessage, sendMsg); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsWebsocketWriteMgsFail, zap.Error(err))
		}
	}
}


```


##### 2.在本项目骨架任意位置，向所有在线的 websocet 客户端广播消息    
> 核心原理：每一个 websocket 客户端都有一个 Hub 结构体，而这个结构体是本项目骨架设置的全局值,因此在任意位置创建一个 websocket 客户端，只要将 Hub 值赋予全局初始化的：variable.WebsocketHub，就可以在任意位置进行广播消息.  
```code   
package demo1

import (
    	serviceWs   "goskeleton/app/service/websocket"
)

// 省略其他无关代码，相关的核心代码如下

if WsHub, ok := variable.WebsocketHub.(*core.Hub); ok {
    // serviceWs 为 app/service/websocket 的别名
   ws := serviceWs.Ws{WsClient: &core.Client{Hub: WsHub}}
   ws.BroadcastMsg("本项目骨架任意位置，使用本段代码对在线的 ws 客户端广播消息")
}

```