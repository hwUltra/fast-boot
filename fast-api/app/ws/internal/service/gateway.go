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
		case 0: //群发
			g.sendToAll(msgBean)
		case 1: //发送给指定用户
			g.sendToUid(msgBean)
		case 2: //发送给群组
			g.sendToGroup(msgBean)
		case 3: //加入群
			g.joinGroup(msgBean)
		case 4: //退出群
			g.leaveGroup(msgBean)
		case 90: //创建群
			g.CreateGroup(uuid.New().String())
		default:
			g.shakeError()
		}
	} else {
		g.shakeError()
	}

}

// 发送给所有人
func (g *Gateway) sendToAll(msgBean *core.MsgBean) {
	g.Ws.SendToAll(core.ReturnBeanSuccess(core.ReturnDataBean{Type: 0, Msg: msgBean.Data}))
}

// 发送给指定用户
func (g *Gateway) sendToUid(msgBean *core.MsgBean) {
	uid, _ := strconv.ParseInt(msgBean.To, 10, 64)
	g.Ws.SendToUid(uid, core.ReturnBeanSuccess(core.ReturnDataBean{
		Type:     1,
		Msg:      msgBean.Data,
		FromPlat: strconv.FormatInt(g.Ws.WsClient.Uid, 10),
		FromUid:  g.Ws.WsClient.Uid,
	}))
}

// 发送给群组
func (g *Gateway) sendToGroup(msgBean *core.MsgBean) {
	g.Ws.SendToGroup(msgBean.To, core.ReturnBeanSuccess(core.ReturnDataBean{
		Type:     2,
		Msg:      msgBean.Data,
		FromPlat: msgBean.To, //群组ID
		FromUid:  g.Ws.WsClient.Uid,
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
func (g *Gateway) CreateGroup(groupId string) {
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
