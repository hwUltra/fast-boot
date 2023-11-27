package service

import (
	"fast-boot/common/websocket/core"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"strconv"
)

type Gateway struct {
	Ws *Ws
}

func NewGateway(ws *Ws) *Gateway {
	return &Gateway{Ws: ws}
}

func (g *Gateway) Handle(data []byte) {

	if msgBean, err := core.GetMsgBean(data); err == nil {
		switch msgBean.Type {
		case "sendAll": //群发
			g.sendToAll(msgBean)
		case "sendUid": //发送给指定用户
			g.sendToUid(msgBean)
		case "sendGroup": //发送给群组
			g.sendToGroup(msgBean)
		case "joinGroup": //加入群
			g.joinGroup(msgBean)
		case "leaveGroup": //退出群
			g.leaveGroup(msgBean)
		case "createGroup": //创建群
			g.createGroup(uuid.New().String())
		default:
			g.shakeError()
		}
	} else {
		g.shakeError()
	}

}

// 发送给所有人
func (g *Gateway) sendToAll(msgBean *core.MsgBean) {
	g.Ws.SendToAll(core.ReturnBeanSuccess(core.ReturnDataBean{
		Type:     "sendAll",
		FromUid:  g.Ws.WsClient.Uid,
		FromPlat: strconv.FormatInt(g.Ws.WsClient.Uid, 10),
		Msg:      msgBean.Data,
	}))
}

// 发送给指定用户
func (g *Gateway) sendToUid(msgBean *core.MsgBean) {
	uid, _ := strconv.ParseInt(msgBean.To, 10, 64)
	g.Ws.SendToUid(uid, core.ReturnBeanSuccess(core.ReturnDataBean{
		Type:     "sendUid",
		FromUid:  g.Ws.WsClient.Uid,
		FromPlat: strconv.FormatInt(g.Ws.WsClient.Uid, 10),
		Msg:      msgBean.Data,
	}))
}

// 发送给群组
func (g *Gateway) sendToGroup(msgBean *core.MsgBean) {
	g.Ws.SendToGroup(msgBean.To, core.ReturnBeanSuccess(core.ReturnDataBean{
		Type:     "sendGroup",
		Msg:      msgBean.Data,
		FromUid:  g.Ws.WsClient.Uid,
		FromPlat: msgBean.To, //群组ID
	}))
}

// 加入群
func (g *Gateway) joinGroup(msgBean *core.MsgBean) {
	err := g.Ws.JoinGroup(g.Ws.WsClient.Uid, msgBean.To)
	msg := core.ReturnBeanMsg("加入成功")
	if err != nil {
		msg = core.ReturnBeanError(err)
	}
	_ = g.Ws.WsClient.SendMessage(websocket.TextMessage, msg)
}

// 退出群 leaveGroup
func (g *Gateway) leaveGroup(msgBean *core.MsgBean) {
	err := g.Ws.LeaveGroup(g.Ws.WsClient.Uid, msgBean.To)
	msg := core.ReturnBeanMsg("退出成功")
	if err != nil {
		msg = core.ReturnBeanError(err)
	}
	_ = g.Ws.WsClient.SendMessage(websocket.TextMessage, msg)
}

// CreateGroup 创建群
func (g *Gateway) createGroup(groupId string) {
	err := g.Ws.CreateGroup(groupId)
	msg := []byte(fmt.Sprintf("创建成功，群组ID：%s", groupId))
	if err != nil {
		msg = core.ReturnBeanError(err)
	}
	_ = g.Ws.WsClient.SendMessage(websocket.TextMessage, msg)
}

// WebsocketHandshakeError
func (g *Gateway) shakeError() {
	if err := g.Ws.WsClient.SendMessage(websocket.TextMessage, []byte(core.WebsocketHandshakeError)); err == nil {
		return
	}
}
