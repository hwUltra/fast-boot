package service

import (
	"encoding/json"
	"fast-boot/common/websocket/core"
	"github.com/gorilla/websocket"
)

type MsgBean struct {
	From string `json:"from,optional"`  //fromId:1=>0, 2=>uid, 3，4，5=>cid
	Type string `json:"type,default=0"` //0.ping  1: 发送广播 2:  发送私聊 3: 发送群聊 4:加入群 5:退出群
	To   string `json:"to,default=0"`   //toId:1=>0, 2=>uid, 3，4，5=>cid
	Msg  string `json:"msg,optional"`
}

type Gateway struct {
	Ws *Ws
}

func NewGateway(ws *Ws) *Gateway {
	return &Gateway{Ws: ws}
}

func (g *Gateway) Handle(data []byte) {
	msgBean := MsgBean{}
	err := json.Unmarshal(data, &msgBean)
	if err != nil {
		if err := g.Ws.WsClient.SendMessage(websocket.TextMessage, []byte(core.WebsocketHandshakeError)); err != nil {
		}
		return
	}
	msgBean.From = g.Ws.WsClient.ClientId

	if msg, err := json.Marshal(&msgBean); err == nil {
		//群发
		g.Ws.SendToAll(msg)
	}

}
