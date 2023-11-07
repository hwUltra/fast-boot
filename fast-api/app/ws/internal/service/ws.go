package service

import (
	"context"
	"fast-boot/app/ws/internal/config"
	"fast-boot/common/websocket/core"
	"fast-boot/common/xerr"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"strconv"
)

type Ws struct {
	WsClient    *core.Client
	RedisClient *redis.Redis
}

// OnOpen 事件函数
func OnOpen(hub *core.Hub, rw http.ResponseWriter, r *http.Request, c config.Config) (*Ws, bool) {
	if client, ok := (&core.Client{}).OnOpen(hub, rw, r); ok {

		ws := &Ws{}
		ws.WsClient = client
		ws.RedisClient = redis.MustNewRedis(c.Redis)

		client.ClientId = uuid.New().String()

		token := r.Header.Get("Sec-WebSocket-Protocol")
		//uid := jwtx.GetUidByToken(c.JWT.AccessSecret, token)
		//上线用上面
		uid, _ := strconv.ParseInt(token, 10, 64)
		fmt.Println("xlsx", token, uid)
		client.Uid = uid
		if uid == 1 {
			client.Role = "admin"
		} else {
			client.Role = "user"
		}
		go client.Heartbeat()

		return ws, ok
	} else {
		return nil, ok
	}
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(context context.Context) {
	go w.WsClient.ReadPump(func(messageType int, receivedData []byte) {

		NewGateway(w).Handle(receivedData)

	}, w.OnError, w.OnClose)
}

// OnError 客户端与服务端在消息交互过程中发生错误回调函数
func (w *Ws) OnError(err error) {
	w.WsClient.State = 0 // 发生错误，状态设置为0, 心跳检测协程则自动退出
	fmt.Printf("clientId:%s远端掉线、卡死、刷新浏览器等会触发该错误: %v\n", w.WsClient.ClientId, err.Error())
}

// OnClose 客户端关闭回调，发生onError回调以后会继续回调该函数
func (w *Ws) OnClose() {

	w.WsClient.Hub.UnRegister <- w.WsClient
}

// GetOnlineClients  获取在线的全部客户端
func (w *Ws) GetOnlineClients() {

	fmt.Printf("在线客户端数量：%d\n", len(w.WsClient.Hub.Clients))
}

// SendToAll  (每一个客户端都有能力)向全部在线客户端广播消息
func (w *Ws) SendToAll(msg []byte) {
	for onlineClient := range w.WsClient.Hub.Clients {
		//获取每一个在线的客户端，向远端发送消息
		if err := onlineClient.SendMessage(websocket.TextMessage, msg); err != nil {
		}
	}
}

// sendToClient 向指定客户端发送消息
func (w *Ws) sendToClient(clientId string, msg []byte) {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.ClientId == clientId {
			if err := onlineClient.SendMessage(websocket.TextMessage, msg); err != nil {
			}
		}
	}
}

// SendToUid 向指定Uid发送消息
func (w *Ws) SendToUid(uid int64, msg []byte) {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.Uid == uid {
			if err := onlineClient.SendMessage(websocket.TextMessage, msg); err != nil {
			}
		}
	}
}

// IsOnline  判断client_id是否还在线
func (w *Ws) IsOnline(clientId string) bool {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.ClientId == clientId {
			return true
		}
	}
	return false
}

// IsUidOnline  判断Uid是否还在线
func (w *Ws) IsUidOnline(uid int64) bool {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.Uid == uid {
			return true
		}
	}
	return false
}

// CloseClient 断开与client_id对应的客户端的连接
func (w *Ws) CloseClient(clientId string) {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.ClientId == clientId {
			onlineClient.Hub.UnRegister <- w.WsClient
		}
	}
}

// CloseUid 断开与client_id对应的客户端的连接
func (w *Ws) CloseUid(uid int64) {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.Uid == uid {
			onlineClient.Hub.UnRegister <- w.WsClient
		}
	}
}

// GetClientByUid 获取Client
func (w *Ws) GetClientByUid(uid int64) (*core.Client, bool) {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.Uid == uid {
			return onlineClient, true
		}
	}
	return nil, false
}

// GetClientById 获取Client
func (w *Ws) GetClientById(clientId string) (*core.Client, bool) {
	for onlineClient := range w.WsClient.Hub.Clients {
		if onlineClient.ClientId == clientId {
			return onlineClient, true
		}
	}
	return nil, false
}

// CreateGroup 创建群组
func (w *Ws) CreateGroup(groupId string) error {
	_, err := w.RedisClient.Sadd("Group", groupId)
	return err
}

// Ungroup 取消群组
func (w *Ws) Ungroup(groupId string) error {
	_, err := w.RedisClient.Srem("Group", groupId)
	return err
}

// JoinGroup 加入群组
func (w *Ws) JoinGroup(uid int64, groupId string) error {
	//
	if ok, _ := w.RedisClient.Sismember("Group", groupId); ok == false {
		return xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "群组不存在")
	}
	if _, err := w.RedisClient.Sadd(groupId, uid); err != nil {
		return xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "加入失败")
	}
	return nil
}

// LeaveGroup 离开群组
func (w *Ws) LeaveGroup(uid int64, groupId string) error {
	if ok, _ := w.RedisClient.Sismember("Group", groupId); ok == false {
		return xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "群组不存在")
	}
	if _, err := w.RedisClient.Srem(groupId, uid); err != nil {
		return xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "退出失败")
	}
	return nil
}

// SendToGroup 向某个分组的所有在线发送数据 ）
func (w *Ws) SendToGroup(groupId string, msg []byte) {
	for onlineClient := range w.WsClient.Hub.Clients {
		if ok, _ := w.RedisClient.Sismember(groupId, onlineClient.Uid); ok {
			_ = onlineClient.SendMessage(websocket.TextMessage, msg)
		}
	}
}
