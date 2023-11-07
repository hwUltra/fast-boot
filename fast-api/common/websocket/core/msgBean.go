package core

import "encoding/json"

// MsgBean 发送
type MsgBean struct {
	Type int8   `json:"type,default=0"` //0: 发送广播 1:  发送私聊 2: 发送群聊 3:加入群 4:退出群
	To   string `json:"to,default=0"`   //uid,cid
	Data string `json:"data,optional"`
}

func GetMsgBean(data []byte) (*MsgBean, error) {
	msgBean := MsgBean{}
	err := json.Unmarshal(data, &msgBean)
	return &msgBean, err
}
