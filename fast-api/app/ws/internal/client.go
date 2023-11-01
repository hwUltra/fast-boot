package internal

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	// send buffer size
	bufSize = 256

	//MaxMessageSize 从消息管道读取消息的最大字节
	MaxMessageSize = 65535

	// HeartbeatFailMaxTime 允许心跳失败的最大次数（默认设置为PingPeriod=30秒检测一次，连续4次没有心跳就会清除后端在线信息）
	HeartbeatFailMaxTime = 4

	WebsocketHandshakeSuccess = `{"code":200,"msg":"ws连接成功","data":""}`

	WebsocketServerPingMsg = "Server->Ping->Client"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	HeartbeatFailTimes int
	Uid                int64
	State              uint8 // ws状态，1=ok；0=出错、掉线等
	sync.RWMutex
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// SendMessage 发送消息，请统一调用本函数进行发送
// 消息发送时增加互斥锁，加强并发情况下程序稳定性
// 提醒：开发者发送消息时，不要调用 c.conn.WriteMessage(messageType, []byte(message)) 直接发送消息
func (c *Client) SendMessage(messageType int, message string) error {
	c.Lock()
	defer func() {
		c.Unlock()
	}()
	// 发送消息时，必须设置本次消息的最大允许时长(秒)
	if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		return err
	}
	if err := c.conn.WriteMessage(messageType, []byte(message)); err != nil {
		return err
	} else {
		return nil
	}
}

// Heartbeat 按照websocket标准协议实现隐式心跳,Server端向Client远端发送ping格式数据包,浏览器收到ping标准格式，自动将消息原路返回给服务器
func (c *Client) Heartbeat() {
	//  1. 设置一个时钟，周期性的向client远端发送心跳数据包
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		err := recover()
		if err != nil {
			//if val, ok := err.(error); ok {
			//	variable.ZapLog.Error(my_errors.ErrorsWebsocketBeatHeartFail, zap.Error(val))
			//}
		}
		ticker.Stop() // 停止该client的心跳检测
	}()
	//2.浏览器收到服务器的ping格式消息，会自动响应pong消息，将服务器消息原路返回过来
	if pongWait == 0 {
		_ = c.conn.SetReadDeadline(time.Time{})
	} else {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	}
	c.conn.SetPongHandler(func(receivedPong string) error {
		if pongWait > time.Nanosecond {
			_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		} else {
			_ = c.conn.SetReadDeadline(time.Time{})
		}
		//fmt.Println("浏览器收到ping标准格式，自动将消息原路返回给服务器：", received_pong)  // 接受到的消息叫做pong，实际上就是服务器发送出去的ping数据包
		return nil
	})
	//3.自动心跳数据
	for {
		select {
		case <-ticker.C:
			if c.State == 1 {
				if err := c.SendMessage(websocket.PingMessage, WebsocketServerPingMsg); err != nil {
					c.HeartbeatFailTimes++
					if c.HeartbeatFailTimes > HeartbeatFailMaxTime {
						c.State = 0
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

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, bufSize),
	}

	if err := client.SendMessage(websocket.TextMessage, WebsocketHandshakeSuccess); err != nil {
		//Log
	}
	// 设置最大读取长度
	client.conn.SetReadLimit(MaxMessageSize)

	client.hub.Register <- client
	client.State = 1

	fmt.Println("r.URL.Query().Get(uid):", r.URL.Query().Get("uid"))
	client.Uid, _ = strconv.ParseInt(r.URL.Query().Get("uid"), 10, 64)

	go client.writePump()
	go client.readPump()
}
