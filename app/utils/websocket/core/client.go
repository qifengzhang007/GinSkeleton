package core

import (
	"errors"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	Hub                *Hub            // 负责处理客户端注册、注销、在线管理
	Conn               *websocket.Conn // 一个ws连接
	Send               chan []byte     // 一个ws连接存储自己的消息管道
	PingPeriod         time.Duration
	ReadDeadline       time.Duration
	WriteDeadline      time.Duration
	HeartbeatFailTimes int
	State              uint8 // ws状态，1=ok；0=出错、掉线等
	sync.RWMutex
}

// 处理握手+协议升级
func (c *Client) OnOpen(context *gin.Context) (*Client, bool) {
	// 1.升级连接,从http--->websocket
	defer func() {
		err := recover()
		if err != nil {
			if val, ok := err.(error); ok {
				variable.ZapLog.Error(my_errors.ErrorsWebsocketOnOpenFail, zap.Error(val))
			}
		}
	}()
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  variable.ConfigYml.GetInt("Websocket.WriteReadBufferSize"),
		WriteBufferSize: variable.ConfigYml.GetInt("Websocket.WriteReadBufferSize"),
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 2.将http协议升级到websocket协议.初始化一个有效的websocket长连接客户端
	if wsConn, err := upGrader.Upgrade(context.Writer, context.Request, nil); err != nil {
		variable.ZapLog.Error(my_errors.ErrorsWebsocketUpgradeFail + err.Error())
		return nil, false
	} else {
		if wsHub, ok := variable.WebsocketHub.(*Hub); ok {
			c.Hub = wsHub
		}
		c.Conn = wsConn
		c.Send = make(chan []byte, variable.ConfigYml.GetInt("Websocket.WriteReadBufferSize"))
		c.PingPeriod = time.Second * variable.ConfigYml.GetDuration("Websocket.PingPeriod")
		c.ReadDeadline = time.Second * variable.ConfigYml.GetDuration("Websocket.ReadDeadline")
		c.WriteDeadline = time.Second * variable.ConfigYml.GetDuration("Websocket.WriteDeadline")

		if err := c.SendMessage(websocket.TextMessage, variable.WebsocketHandshakeSuccess); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsWebsocketWriteMgsFail, zap.Error(err))
		}
		c.Conn.SetReadLimit(variable.ConfigYml.GetInt64("Websocket.MaxMessageSize")) // 设置最大读取长度
		c.Hub.Register <- c
		c.State = 1
		return c, true
	}

}

// 主要功能主要是实时接收消息
func (c *Client) ReadPump(callbackOnMessage func(messageType int, receivedData []byte), callbackOnError func(err error), callbackOnClose func()) {
	// 回调 onclose 事件
	defer func() {
		err := recover()
		if err != nil {
			if realErr, isOk := err.(error); isOk {
				variable.ZapLog.Error(my_errors.ErrorsWebsocketReadMessageFail, zap.Error(realErr))
			}
		}
		callbackOnClose()
	}()

	// OnMessage事件
	for {
		if c.State == 1 {
			mt, bReceivedData, err := c.Conn.ReadMessage()
			if err == nil {
				callbackOnMessage(mt, bReceivedData)
			} else {
				// OnError事件读（消息出错)
				callbackOnError(err)
				break
			}
		} else {
			// OnError事件(状态不可用，一般是程序事先检测到双方无法进行通信，进行的回调)
			callbackOnError(errors.New(my_errors.ErrorsWebsocketStateInvalid))
			break
		}

	}
}

// 发送消息，请统一调用本函数进行发送
// 消息发送时增加互斥锁，加强并发情况下程序稳定性
// 提醒：开发者发送消息时，不要调用 c.Conn.WriteMessage(messageType, []byte(message)) 直接发送消息
func (c *Client) SendMessage(messageType int, message string) error {
	c.Lock()
	defer func() {
		c.Unlock()
	}()
	// 发送消息时，必须设置本次消息的最大允许时长(秒)
	if err := c.Conn.SetWriteDeadline(time.Now().Add(c.WriteDeadline)); err != nil {
		variable.ZapLog.Error(my_errors.ErrorsWebsocketSetWriteDeadlineFail, zap.Error(err))
		return err
	}
	if err := c.Conn.WriteMessage(messageType, []byte(message)); err != nil {
		return err
	} else {
		return nil
	}
}

// 按照websocket标准协议实现隐式心跳,Server端向Client远端发送ping格式数据包,浏览器收到ping标准格式，自动将消息原路返回给服务器
func (c *Client) Heartbeat() {
	//  1. 设置一个时钟，周期性的向client远端发送心跳数据包
	ticker := time.NewTicker(c.PingPeriod)
	defer func() {
		err := recover()
		if err != nil {
			if val, ok := err.(error); ok {
				variable.ZapLog.Error(my_errors.ErrorsWebsocketBeatHeartFail, zap.Error(val))
			}
		}
		ticker.Stop() // 停止该client的心跳检测
	}()
	//2.浏览器收到服务器的ping格式消息，会自动响应pong消息，将服务器消息原路返回过来
	if c.ReadDeadline == 0 {
		_ = c.Conn.SetReadDeadline(time.Time{})
	} else {
		_ = c.Conn.SetReadDeadline(time.Now().Add(c.ReadDeadline))
	}
	c.Conn.SetPongHandler(func(receivedPong string) error {
		if c.ReadDeadline > time.Nanosecond {
			_ = c.Conn.SetReadDeadline(time.Now().Add(c.ReadDeadline))
		} else {
			_ = c.Conn.SetReadDeadline(time.Time{})
		}
		//fmt.Println("浏览器收到ping标准格式，自动将消息原路返回给服务器：", received_pong)  // 接受到的消息叫做pong，实际上就是服务器发送出去的ping数据包
		return nil
	})
	//3.自动心跳数据
	for {
		select {
		case <-ticker.C:
			if c.State == 1 {
				if err := c.SendMessage(websocket.PingMessage, variable.WebsocketServerPingMsg); err != nil {
					c.HeartbeatFailTimes++
					if c.HeartbeatFailTimes > variable.ConfigYml.GetInt("Websocket.HeartbeatFailMaxTimes") {
						c.State = 0
						variable.ZapLog.Error(my_errors.ErrorsWebsocketBeatHeartsMoreThanMaxTimes, zap.Error(err))
						return
					}
				} else {
					if c.HeartbeatFailTimes > 0 {
						c.HeartbeatFailTimes--
					}
				}
			} else {
				return
			}

		}
	}
}
