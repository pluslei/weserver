package mqtt

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/astaxie/beego"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type message struct {
	topic   string
	message interface{}
}

type MQ struct {
	topic  map[string]byte //订阅列表
	qos    byte
	retain bool
	opts   *MQTT.ClientOptions
	conn   MQTT.Client
	msgch  chan *message
}

func NewMQ(o *Configer) *MQ {
	m := MQ{}
	opts := MQTT.NewClientOptions().AddBroker(o.MqAddress)
	opts.SetUsername(o.MqUserName)
	opts.SetPassword(o.MqPwd)
	opts.SetClientID(o.MqClientID)
	opts.SetCleanSession(o.MqIsCleansession)
	opts.SetProtocolVersion(uint(o.MqVersion))
	opts.SetAutoReconnect(o.MqIsreconnect) //自动重连
	opts.SetDefaultPublishHandler(messageHandler)
	m.topic = o.MqTopic
	m.opts = opts
	m.msgch = make(chan *message, 102400)
	return &m
}

// callback 异步接收消息
func messageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func (m *MQ) Runing() {
	err := m.connect(m.topic)
	if err != nil {
		beego.Error("mqtt connect: ", err)
	} else {
		beego.Info("mqtt running OK ")
		go m.worker()
	}
}

//发送消息线程
func (m *MQ) worker() {
	for {
		msg, ok := <-m.msgch
		if ok {
			m.qos = 0
			m.retain = false
			tokenSen := m.conn.Publish(msg.topic, m.qos, m.retain, msg.message)
			tokenSen.Wait()
			if tokenSen.Error() != nil {
				beego.Debug("publish error:", tokenSen.Error())
			}
		} else {
			beego.Error("mqtt publish worker shutdown!!! ")
		}
	}
}

//go 实现 Hmac-SHA1
func (m *MQ) goHmacSHA1(id, key string) string {
	//sha1
	h := sha1.New()
	io.WriteString(h, id)
	//hmac ,use sha1
	newkey := []byte(key)
	mac := hmac.New(sha1.New, newkey)
	mac.Write([]byte(id))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

// 连接并订阅消息
func (m *MQ) connect(topic map[string]byte) error {
	m.conn = MQTT.NewClient(m.opts)
	if tokenConn := m.conn.Connect(); tokenConn.Wait() && tokenConn.Error() != nil {
		beego.Error("connect error", tokenConn.Error())
		return tokenConn.Error()
	}
	/* 单聊天室
	if tokenSub := m.conn.Subscribe(topic, 1, nil); tokenSub.Wait() && tokenSub.Error() != nil {
		beego.Error("topic sub error", tokenSub.Error())
		return tokenSub.Error()
	}
	*/
	//多聊天室
	if tokenSub := m.conn.SubscribeMultiple(topic, nil); tokenSub.Wait() && tokenSub.Error() != nil {
		beego.Error("topic sub error", tokenSub.Error())
		return tokenSub.Error()
	}
	return nil
}

// 取消订阅
func (m *MQ) UnSubScribe(topic string) error {
	if tokenUnS := m.conn.Unsubscribe(topic); tokenUnS.Wait() && tokenUnS.Error() != nil {
		beego.Error("topic Unsubscribe error: ", tokenUnS.Error())
		return tokenUnS.Error()
	}
	return nil
}

// 断开连接
func (m *MQ) DisConnect(client MQTT.Client, quiesce uint) {
	m.conn.Disconnect(quiesce)
	m.conn = nil
}

//发消息
func (m *MQ) sendMessage(topic string, args interface{}) error {
	_, ok := m.topic[topic]
	if !ok {
		beego.Error("Topic is no permission")
		return fmt.Errorf("Topic is no persission")
	}
	msg := &message{
		topic:   topic,
		message: args,
	}
	select {
	case m.msgch <- msg:
		return nil
	default:
		beego.Error("message ch full")
		return fmt.Errorf("message ch full")
	}
}
