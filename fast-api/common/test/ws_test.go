package test

import (
	"encoding/json"
	wsTools "github.com/hwUltra/fb-tools/websocket"
	"testing"
)

func Test_Ws_SendMsg(t *testing.T) {
	host := "127.0.0.1:7770"
	path := "ws"
	query := "uid=88"
	token := "33"

	dataByte, _ := json.Marshal(map[string]interface{}{
		"type": "sendAll",
		"to":   "all",
		"data": "success",
	})
	err := wsTools.SendMsg(host, path, query, token, dataByte)
	if err != nil {
		t.Errorf("wsTools SendMsg write: %v", err)
	}
}
