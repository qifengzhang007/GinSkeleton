package Core

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 5 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// 1.将http连接升级到websocket连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024 * 2,
	WriteBufferSize: 65535,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from th
//
//
//
//s goroutine.
// 主要功能是接收消息
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	//服务器发送Ping格式消息，浏览器自动响应pong，
	c.Conn.SetPongHandler(func(pong string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		fmt.Println("服务端收到了客户端的pong消息", pong)
		return nil
	})

	for {
		// c.Conn 就是一个有效的websocket客户端，就会收到来自网页客户端的消息，这里读取消息
		// 如果远端断开了，name读取消息就返回错误
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("读取管道发生错误: %v", err.Error())
			//捕获远端可能断开的错误（是否存在未知的关闭错误）
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.

// 核心功能是从Client的send 管道读取消息，发送给网页远端客户端（例如网页客户端）
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				// 当client的消息管道无法读取消息（可能由于关闭等原因引起，那么我们就先关闭“消息写入器”，然后出发defer函数，关闭websocket连接）
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//NextWriter 函数获取一个消息写入器（用于发送消息使用，一个连接最多只能有一个消息写入器）
			/*			w, err := c.Conn.NextWriter(websocket.TextMessage)
						if err != nil {
							return
						}
						//向远端（例如网页客户端）发送消息
						w.Write(message)*/
			c.Conn.WriteMessage(websocket.TextMessage, message) // 向远端发送消息
			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				//w.Write(newline)
				//w.Write(<-c.Send)
				c.Conn.WriteMessage(websocket.TextMessage, newline)
				c.Conn.WriteMessage(websocket.TextMessage, <-c.Send)
			}

			//if err := w.Close(); err != nil {
			//	return
			//}
			//websocket ping pong:目前的话，浏览器中没有相关api发送ping给服务器，只能由服务器发ping给浏览器，浏览器返回pong消息；
		case tmp := <-ticker.C:
			fmt.Println(tmp)
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte("Server_Ping")); err != nil {
				fmt.Println("ping消息发送发生错误，于客户端通信受阻" + err.Error())
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 2.将http协议升级到websocket协议.初始化一个有效的websocket长连接客户端
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	//websocket客户端上线就注册在hub中心
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	// 启动新的协程独立处理通道中的消息
	go client.ReadPump()
	go client.WritePump()
}
