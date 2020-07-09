package Core

import (
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/Config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Client struct {
	Hub                *Hub            // 负责处理客户端注册、注销、在线管理
	Conn               *websocket.Conn // 一个ws连接
	Send               chan []byte     // 一个ws连接存储自己的消息管道
	PingPeriod         time.Duration
	PongWait           time.Duration
	ReadDeadline       time.Duration
	HeartbeatFailTimes int
}

// 处理握手+协议升级
func (c *Client) OnOpen(context *gin.Context) (*Client, bool) {
	// 1.升级连接,从http--->websocket

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  Config.CreateYamlFactory().GetInt("Websocket.WriteReadBufferSize"),
		WriteBufferSize: Config.CreateYamlFactory().GetInt("Websocket.WriteReadBufferSize"),
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 2.将http协议升级到websocket协议.初始化一个有效的websocket长连接客户端
	if ws_conn, err := upgrader.Upgrade(context.Writer, context.Request, nil); err != nil {
		log.Println(MyErrors.Errors_Websocket_OnOpen_Fail, err.Error())
		return nil, false
	} else {
		if ws_hub, ok := Variable.Websocket_Hub.(*Hub); ok {
			c.Hub = ws_hub
		}
		c.Conn = ws_conn
		c.Send = make(chan []byte, Config.CreateYamlFactory().GetInt("Websocket.WriteReadBufferSize"))
		c.PingPeriod = Config.CreateYamlFactory().GetDuration("Websocket.PingPeriod")
		c.PongWait = Config.CreateYamlFactory().GetDuration("Websocket.PingPeriod") + 3 //  官方文档中注明，pongwait 必须> ping 时间
		c.ReadDeadline = Config.CreateYamlFactory().GetDuration("Websocket.ReadDeadline")
		c.Hub.Register <- c
		ws_conn.WriteMessage(websocket.TextMessage, []byte(Variable.Websocket_Handshake_Success))
		return c, true
	}

}

// 主要功能主要是实时接收消息
func (c *Client) ReadPump(callback_on_message func(message_type int, received_data []byte), callback_on_error func(err error), callback_on_close func()) {

	// 回调 onclose 事件
	defer func() {
		callback_on_close()
	}()
	// OnMessage事件
	c.Conn.SetReadLimit(Config.CreateYamlFactory().GetInt64("Websocket.MaxMessageSize")) // 设置最大读取长度
	for {
		mt, b_received_data, err := c.Conn.ReadMessage()
		if err == nil {
			callback_on_message(mt, b_received_data)
		} else {
			// OnError事件
			callback_on_error(err)
			break
		}
		if c.ReadDeadline > 0 {
			c.Conn.SetReadDeadline(time.Now().Add(c.ReadDeadline * time.Second))
		}
	}
}

// 按照websocket标准协议实现隐式心跳,Server端向Client远端发送ping格式数据包,浏览器收到ping标准格式，自动将消息原路返回给服务器
func (c *Client) Heartbeat(callback_close func()) {

	//2.浏览器收到服务器的ping格式消息，会自动响应pong消息，将服务器消息原路返回过来
	c.Conn.SetPongHandler(func(pong string) error {
		//fmt.Println("浏览器收到ping标准格式，自动将消息原路返回给服务器：", pong)  // 接受到的消息叫做pong，实际上就是服务器发送出去的ping数据包
		return nil
	})

	//  1. 设置一个时钟，周期性的向client远端发送心跳数据包
	ticker := time.NewTicker(c.PingPeriod * time.Second)
	defer func() {
		ticker.Stop()    // 停止该client的心跳检测
		callback_close() // 注销 client
	}()

	for {
		select {
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte(Variable.Websocket_Server_Ping_Msg)); err != nil {
				c.HeartbeatFailTimes++
				if c.HeartbeatFailTimes > Config.CreateYamlFactory().GetInt("Websocket.HeartbeatFailMaxTimes") {
					return
				}
			} else {
				if c.HeartbeatFailTimes > 0 {
					c.HeartbeatFailTimes--
				}
			}
		}
	}
}
