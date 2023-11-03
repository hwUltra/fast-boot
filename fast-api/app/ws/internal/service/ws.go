package service

import (
	"context"
	"fast-boot/common/websocket/core"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
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
func OnOpen(hub *core.Hub, rw http.ResponseWriter, r *http.Request) (*Ws, bool) {
	if client, ok := (&core.Client{}).OnOpen(hub, rw, r); ok {
		go client.Heartbeat()
		ws := &Ws{}
		ws.WsClient = client
		return ws, true
	} else {
		return nil, false
	}
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(context context.Context) {
	go w.WsClient.ReadPump(func(messageType int, receivedData []byte) {
		//参数说明
		//messageType 消息类型，1=文本
		//receivedData 服务器接收到客户端（例如js客户端）发来的的数据，[]byte 格式
		tempMsg := "服务器已经收到了你的消息==>" + string(receivedData)
		fmt.Println("OnMessage", tempMsg, messageType)
		// 回复客户端已经收到消息;
		//if err := w.WsClient.SendMessage(messageType, tempMsg); err != nil {
		//	fmt.Println("消息发送出现错误", err)
		//}

		//给全部用户发送
		for onlineClient := range w.WsClient.Hub.Clients {
			//获取每一个在线的客户端，向远端发送消息
			if err := onlineClient.SendMessage(messageType, tempMsg); err != nil {
			}
		}

		fmt.Printf("在线客户端数量：%d\n", len(w.WsClient.Hub.Clients))

	}, w.OnError, w.OnClose)
}

// OnError 客户端与服务端在消息交互过程中发生错误回调函数
func (w *Ws) OnError(err error) {
	w.WsClient.State = 0 // 发生错误，状态设置为0, 心跳检测协程则自动退出
	fmt.Printf("远端掉线、卡死、刷新浏览器等会触发该错误: %v\n", err.Error())
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

		}
	}
}
