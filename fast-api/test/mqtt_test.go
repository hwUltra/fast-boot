package test

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

func TestMqtt(t *testing.T) {
	var broker = "127.0.0.1" //地址
	var port = 1883          // 端口
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client111") //设备唯一id，正常应该是设备拿自己的设备id注册到服务器上。
	opts.SetUsername("emqx")              //账号
	opts.SetPassword("public")            //密码
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)     //等待消息
	publish(client) //发消息

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil) ////这里如果不指定方法，就用的上面的SetDefaultPublishHandler设置的方法
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
