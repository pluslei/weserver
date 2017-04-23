package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/astaxie/beego"
	// for json get
	"time"
)

var Lock sync.Mutex

type Wechat struct {
	appID                string
	appSecret            string
	accessTokenFetchUrl  string
	customServicePostUrl string
	AccessToken          string  `json:"access_token"`
	ExpiresIn            float64 `json:"expires_in"`
	Errcode              float64 `json:"errcode"`
	Errmsg               string  `json:"errmsg"`
	TextUrl              string
	msgch                chan []byte
}

// text msg
type CustomMsg struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	Text    TextMsgContent `json:"text"`
}

type TextMsgContent struct {
	Content string `json:"content"`
}

func Start(p *Config) *Wechat {
	w := Wechat{}
	w.appID = p.appID
	w.appSecret = p.appSecret
	w.accessTokenFetchUrl = p.accessTokenFetchUrl
	w.customServicePostUrl = p.customServicePostUrl
	w.msgch = make(chan []byte, 102400)
	return &w
}

func (w *Wechat) Work() {
	Lock.Lock()
	go w.getAccessToken()
	Lock.Unlock()
	time.Sleep(time.Second * 1)
	go w.send()

}

func (w *Wechat) getAccessToken() (string, float64, error) {
	t := time.NewTimer(time.Hour * 1)
	for {
		requestLine := fmt.Sprintf(w.accessTokenFetchUrl, w.appID, w.appSecret)
		beego.Debug("GetTokenString", requestLine)

		resp, err := http.Get(requestLine)
		if err != nil || resp.StatusCode != http.StatusOK {
			return "", 0.0, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", 0.0, err
		}
		beego.Debug(string(body))

		if bytes.Contains(body, []byte("access_token")) {
			beego.Debug("Request ok!!!!!!!")
			err = json.Unmarshal(body, w)
			if err != nil {
				return "", 0.0, err
			}
			w.TextUrl = fmt.Sprintf(w.customServicePostUrl, w.AccessToken)
		} else {
			beego.Debug("Request error!!!!!!!")
			err = json.Unmarshal(body, w)
			if err != nil {
				return "", 0.0, err
			}
		}
		beego.Debug(w.AccessToken, w.ExpiresIn, w.TextUrl)
		<-t.C
	}
}

func (w *Wechat) send() {
	for {
		msg, ok := <-w.msgch
		if ok {
			postReq, err := http.NewRequest("POST", w.TextUrl, bytes.NewReader(msg))
			if err != nil {
				beego.Debug("POST WeChatText fail", err)
			}
			postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

			client := &http.Client{}
			resp, err := client.Do(postReq)
			if err != nil {
				beego.Debug("client.Do() WeChatText fail", err)
			}
			resp.Body.Close()
		} else {
			beego.Error("Wechat Publish msg shutdown!!! ")
		}
	}
}

func (w *Wechat) sendCustomTxTMsg(openId, msg string) error {

	TxtMsg := &CustomMsg{
		ToUser:  openId,
		MsgType: "text",
		Text:    TextMsgContent{Content: msg},
	}
	body, err := json.MarshalIndent(TxtMsg, " ", "  ")
	if err != nil {
		return err
	}
	select {
	case w.msgch <- body:
		return nil
	default:
		beego.Error("wechat message ch full")
		return fmt.Errorf("wechat message ch full")
	}
}
