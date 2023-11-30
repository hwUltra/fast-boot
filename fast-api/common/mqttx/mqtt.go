package mqttx

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttX struct {
	Conf   Conf
	Client mqtt.Client
}

type MsgHandler func(MqtMsg)
type OnConnectHandler func()
type OnConnectionLostHandler func(error)

func Create(conf Conf, clientID string,
	defaultPublishHandler MsgHandler,
	onConnectHandler OnConnectHandler,
	onConnectionLostHandler OnConnectionLostHandler,
) (*MqttX, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", conf.Host, conf.Port))
	opts.SetUsername(conf.Username)
	opts.SetPassword(conf.Password)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		if defaultPublishHandler != nil {
			defaultPublishHandler(MqtMsg{
				Topic:     message.Topic(),
				Duplicate: message.Duplicate(),
				Qos:       message.Qos(),
				Retained:  message.Retained(),
				MessageID: message.MessageID(),
				Payload:   message.Payload(),
				Ack:       message.Ack,
			})
		}
	})
	opts.OnConnect = func(client mqtt.Client) {
		if onConnectHandler != nil {
			onConnectHandler()
		}
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		if onConnectionLostHandler != nil {
			onConnectionLostHandler(err)
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MqttX{
		Conf:   conf,
		Client: client,
	}, nil
}

func QuickCreate(conf Conf, clientID string) (*MqttX, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", conf.Host, conf.Port))
	opts.SetUsername(conf.Username)
	opts.SetPassword(conf.Password)
	opts.SetClientID(clientID)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &MqttX{
		Conf:   conf,
		Client: client,
	}, nil
}

func (mx *MqttX) Publish(msg MqtMsg) {
	//client mqtt.Client
	token := mx.Client.Publish(msg.Topic, msg.Qos, msg.Retained, msg.Payload)
	token.Wait()
}

func (mx *MqttX) Sub(msg MqtMsg, callback MsgHandler) {
	//这里如果不指定方法，就用的上面的SetDefaultPublishHandler设置的方法
	var token mqtt.Token
	if callback != nil {
		token = mx.Client.Subscribe(msg.Topic, msg.Qos, func(client mqtt.Client, message mqtt.Message) {
			callback(MqtMsg{
				Topic:     message.Topic(),
				Duplicate: message.Duplicate(),
				Qos:       message.Qos(),
				Retained:  message.Retained(),
				MessageID: message.MessageID(),
				Payload:   message.Payload(),
				Ack:       message.Ack,
			})
		})
	} else {
		token = mx.Client.Subscribe(msg.Topic, msg.Qos, nil)
	}
	token.Wait()
}

func (mx *MqttX) Disconnect(quiesce uint) {
	mx.Client.Disconnect(quiesce)
}
