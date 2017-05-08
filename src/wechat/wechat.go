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
	templatePostUrl      string

	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
	Errcode     float64 `json:"errcode"`
	Errmsg      string  `json:"errmsg"`
	TextUrl     string  //post custom msg
	TemplateUrl string  // post template msg
	msgch       chan []byte
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
	w.templatePostUrl = p.templatePostUrl
	w.msgch = make(chan []byte, 102400)
	return &w
}

func (w *Wechat) Running() {
	var status bool = false
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for {
			err := w.getAccessToken()
			if err != nil {
				beego.Error("getAccessToken error: ", err)
				status = true
			}
			if status {
				time.Sleep(10)
			} else {
				break
			}
		}
		for {
			select {
			case msg, ok := <-w.msgch:
				if ok {
					postReq, err := http.NewRequest("POST", w.TextUrl, bytes.NewReader(msg))
					if err != nil {
						beego.Debug("POST WeChatText Fail", err)
						break
					}
					postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

					client := &http.Client{}
					resp, err := client.Do(postReq)
					if resp.StatusCode != http.StatusOK || err != nil {
						beego.Debug("resp.StatusCode error: ", resp.StatusCode, err.Error())
						resp.Body.Close()
						break
					}
					if resp != nil {
						buf, err := ioutil.ReadAll(resp.Body)
						if err != nil {
							beego.Debug("Read Body error", err)
							resp.Body.Close()
							break
						}
						err = json.Unmarshal(buf, w)
						if err != nil {
							if w.Errcode != 0 && w.Errmsg != "ok" {
								beego.Debug("WeChat Error CodeInfo:", w.Errcode, w.Errmsg)
							}
							resp.Body.Close()
							break
						}
						resp.Body.Close()
						break
					}
					resp.Body.Close()
				} else {
					beego.Error("Wechat Publish msg shutdown!!! ")
				}
				break
			case <-ticker.C:
				for {
					err := w.getAccessToken()
					if err != nil {
						beego.Error("getAccessToken error: ", err)
						status = true
					}
					if status {
						time.Sleep(10)
					} else {
						break
					}
				}
				break
			}
		}
	}()
}

func (w *Wechat) getAccessToken() error {

	requestLine := fmt.Sprintf(w.accessTokenFetchUrl, w.appID, w.appSecret)
	beego.Debug("GetTokenString", requestLine)

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	beego.Debug("AccessToken:", string(body))
	if bytes.Contains(body, []byte("access_token")) {
		beego.Debug("Request ok!!!!!!!")
		err = json.Unmarshal(body, w)
		if err != nil {
			return err
		}
		w.TextUrl = fmt.Sprintf(w.customServicePostUrl, w.AccessToken)
		w.TemplateUrl = fmt.Sprintf(w.templatePostUrl, w.AccessToken)
	} else {
		beego.Debug("Request error!!!!!!!")
		err = json.Unmarshal(body, w)
		if err != nil {
			return err
		}
	}
	// beego.Debug(w.AccessToken, w.ExpiresIn, w.TextUrl)
	return nil
}

func (w *Wechat) sendCustomTxTMsg(openId, msg string) error {
	sendMsg := fmt.Sprintf("今日策略：%s\r\nUrl: test.780.com.cn", msg)
	TxtMsg := &CustomMsg{
		ToUser:  openId,
		MsgType: "text",
		Text:    TextMsgContent{Content: sendMsg},
	}
	body, err := json.MarshalIndent(TxtMsg, " ", "  ")
	if err != nil {
		beego.Debug("sendCustomTxTMsg() error:", err)
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
