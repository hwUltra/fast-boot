package test

import (
	"fast-boot/common/mqttx"
	"fmt"
	"testing"
	"time"
)

func TestMQTTX(t *testing.T) {
	mt, err := mqttx.Create(mqttx.Conf{
		Host:     "127.0.0.1",
		Port:     1883,
		Username: "user",
		Password: "123123",
	}, "go_mqtt_client111", func(msg mqttx.MqtMsg) {
		fmt.Printf("defalut Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	}, func() {
		fmt.Println("on connection")
	}, func(err error) {
		fmt.Println("lost connection err", err)
	})
	if err != nil {
		t.Errorf("mqttx CreateMqttX: %v", err)
	}
	topic := "topic/test"
	mt.Sub(mqttx.MqtMsg{Topic: topic, Qos: 1}, func(msg mqttx.MqtMsg) {
		fmt.Printf("TestMQTTX Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	})
	//mt.Sub(mqttx.MqtMsg{Topic: topic, Qos: 1}, nil)

	for i := 0; i < 4; i++ {
		text := fmt.Sprintf("Message %d", i)
		mt.Publish(mqttx.MqtMsg{Topic: topic, Payload: []byte(text)})
		time.Sleep(time.Second)
	}

	mt.Disconnect(250)
}

func TestMQTTXSal(t *testing.T) {
	mt, err := mqttx.QuickCreate(mqttx.Conf{
		Host:     "127.0.0.1",
		Port:     1883,
		Username: "user",
		Password: "123123",
	}, "go_mqtt_client111")
	if err != nil {
		t.Errorf("Quick CreateMqttX: %v", err)
	}
	topic := "topic/test"
	mt.Sub(mqttx.MqtMsg{Topic: topic, Qos: 1}, func(msg mqttx.MqtMsg) {
		fmt.Printf("Quick Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	})

	for i := 0; i < 4; i++ {
		text := fmt.Sprintf("Quick Message %d", i)
		mt.Publish(mqttx.MqtMsg{Topic: topic, Payload: []byte(text)})
		time.Sleep(time.Second)
	}

	mt.Disconnect(250)
}
