package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/astaxie/beego"
	// for json get
)

var Lock sync.Mutex

type SMS struct {
	URL   string
	msgch chan []byte
}

func Start(p *Config) *SMS {
	s := SMS{}
	s.URL = p.Url
	s.msgch = make(chan []byte, 10240)
	return &s
}

func (s *SMS) Running() {
	go func() {
		for {
			msg, ok := <-w.msgch
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
		}
	}()
}

func (s *SMS) sendSMSmsg(phoneNum, msg, sign string) error {
	body := fmt.Sprintf(s.URL, phoneNum, msg, sign)
	select {
	case s.msgch <- body:
		return nil
	default:
		beego.Error("wechat message ch full")
		return fmt.Errorf("wechat message ch full")
	}
}
