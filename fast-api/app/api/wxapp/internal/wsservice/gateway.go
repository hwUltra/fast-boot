package wsservice

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/hwUltra/fb-tools/utils"
	"github.com/hwUltra/fb-tools/wsCore"
	"strconv"
)

type Gateway struct {
	Ws *Ws
}

type ReturnDataBean struct {
	Type      string `json:"type,default=sendAll"`
	FromType  string `json:"fromType,default=0,optional"`  //uid,cid
	FromUid   int64  `json:"fromUid,default=0,optional"`   //谁发的
	FromGroup int64  `json:"fromGroup,default=0,optional"` //谁发的
	Msg       string `json:"msg"`
}

func NewGateway(ws *Ws) *Gateway {
	return &Gateway{Ws: ws}
}

func (g *Gateway) Handle(data []byte) {

	if msgBean, err := wsCore.GetMsgBean(data); err == nil {
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
func (g *Gateway) sendToAll(msgBean *wsCore.MsgBean) {
	if g.Ws.WsClient.Role != "admin" {
		msg := wsCore.ReturnBeanFail("你没有权限发送群消息")
		_ = g.Ws.WsClient.SendMessage(websocket.TextMessage, msg)
		return
	}
	g.Ws.SendToAll(wsCore.ReturnBeanSuccess(ReturnDataBean{
		Type:      "sendAll",
		FromType:  "user",
		FromUid:   g.Ws.WsClient.Uid,
		FromGroup: 0,
		Msg:       msgBean.Data,
	}))
}

// 发送给指定用户
func (g *Gateway) sendToUid(msgBean *wsCore.MsgBean) {
	uid, _ := strconv.ParseInt(msgBean.To, 10, 64)
	g.Ws.SendToUid(uid, wsCore.ReturnBeanSuccess(ReturnDataBean{
		Type:      "sendUid",
		FromType:  "user",
		FromUid:   g.Ws.WsClient.Uid,
		FromGroup: 0,
		Msg:       msgBean.Data,
	}))
}

// 发送给群组
func (g *Gateway) sendToGroup(msgBean *wsCore.MsgBean) {
	g.Ws.SendToGroup(msgBean.To, wsCore.ReturnBeanSuccess(ReturnDataBean{
		Type:      "sendGroup",
		FromType:  "Group",
		Msg:       msgBean.Data,
		FromUid:   g.Ws.WsClient.Uid,
		FromGroup: utils.ToInt(msgBean.To),
	}))
}

// 加入群
func (g *Gateway) joinGroup(msgBean *wsCore.MsgBean) {
	err := g.Ws.JoinGroup(g.Ws.WsClient.Uid, msgBean.To)
	msg := wsCore.ReturnBeanMsg("加入成功")
	if err != nil {
		msg = wsCore.ReturnBeanError(err)
	}
	_ = g.Ws.WsClient.SendMessage(websocket.TextMessage, msg)
}

// 退出群 leaveGroup
func (g *Gateway) leaveGroup(msgBean *wsCore.MsgBean) {
	err := g.Ws.LeaveGroup(g.Ws.WsClient.Uid, msgBean.To)
	msg := wsCore.ReturnBeanMsg("退出成功")
	if err != nil {
		msg = wsCore.ReturnBeanError(err)
	}
	_ = g.Ws.WsClient.SendMessage(websocket.TextMessage, msg)
}

// CreateGroup 创建群
func (g *Gateway) createGroup(groupId string) {
	err := g.Ws.CreateGroup(groupId)
	msg := []byte(fmt.Sprintf("创建成功，群组ID：%s", groupId))
	if err != nil {
		msg = wsCore.ReturnBeanError(err)
	}
	_ = g.Ws.WsClient.SendMessage(websocket.TextMessage, msg)
}

// WebsocketHandshakeError
func (g *Gateway) shakeError() {
	if err := g.Ws.WsClient.SendMessage(websocket.TextMessage, []byte(wsCore.WebsocketHandshakeError)); err == nil {
		return
	}
}
