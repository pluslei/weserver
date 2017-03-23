package mqtt

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/astaxie/beego"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// callback 异步接收消息
var FU MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

//go 实现 Hmac-SHA1
func goHmacSHA1(id, key string) string {
	//sha1
	h := sha1.New()
	io.WriteString(h, id)
	//hmac ,use sha1
	newkey := []byte(key)
	mac := hmac.New(sha1.New, newkey)
	mac.Write([]byte(id))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

// 设置连接参数
func SetConnectOptions(address, userName, pwd, clientID string, version uint, reconnect, clean bool, f MQTT.MessageHandler) *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions().AddBroker(address)
	opts.SetUsername(userName)
	opts.SetPassword(pwd)
	opts.SetClientID(clientID)
	opts.SetCleanSession(clean)
	opts.SetProtocolVersion(version)
	opts.SetAutoReconnect(reconnect) //自动重连
	opts.SetDefaultPublishHandler(f)
	return opts
}

// 连接并订阅消息
func ConnectSubScribe(opts *MQTT.ClientOptions, topic string) MQTT.Client {
	client := MQTT.NewClient(opts)
	if tokenConn := client.Connect(); tokenConn.Wait() && tokenConn.Error() != nil {
		beego.Error("connect error", tokenConn.Error())
		panic(tokenConn.Error())
	}
	if tokenSub := client.Subscribe(topic, 1, nil); tokenSub.Wait() && tokenSub.Error() != nil {
		beego.Error("topic sub error", tokenSub.Error())
		panic(tokenSub.Error())
	}
	return client
}

// 取消订阅
func UnSubScribe(client MQTT.Client, topic string) {
	if tokenUnS := client.Unsubscribe(topic); tokenUnS.Wait() && tokenUnS.Error() != nil {
		panic(tokenUnS.Error())
	}
}

// 断开连接
func DisConnect(client MQTT.Client, quiesce uint) {
	client.Disconnect(quiesce)
}

//发消息
func SendMessage(topic string, message interface{}, qos byte, retain bool, client MQTT.Client) {
	tokenSen := client.Publish(topic, qos, retain, message)
	tokenSen.Wait()
	if tokenSen.Error() != nil {
		panic(tokenSen.Error)
	}
}
