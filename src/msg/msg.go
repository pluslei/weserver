package msg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/astaxie/beego"
	// for json get
)

var Lock sync.Mutex

type SMS struct {
	URL      string
	USER_URL string
	msgch    chan string
}

func Start(p *Config) *SMS {
	s := SMS{}
	s.URL = p.Url
	s.USER_URL = p.USER_Url
	s.msgch = make(chan string, 10240)
	return &s
}

func (s *SMS) Running() {
	go func() {
		for {
			url, ok := <-s.msgch
			beego.Debug("url", url)
			if ok {
				postReq, err := http.NewRequest("POST", url, strings.NewReader("name=PostName"))
				if err != nil {
					beego.Debug("POST WeChatText Fail", err)
					break
				}
				postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				client := &http.Client{}
				resp, err := client.Do(postReq)
				if resp.StatusCode != http.StatusOK || err != nil {
					beego.Debug("resp.Status error: ", resp.Status, err)
					resp.Body.Close()
					break
				}
				if resp != nil {
					_, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						beego.Debug("Read Body error", err)
						resp.Body.Close()
						break
					}
					beego.Debug("resp code", resp)
					// err = json.Unmarshal(buf, w)
					// if err != nil {
					// 	if w.Errcode != 0 && w.Errmsg != "ok" {
					// 		beego.Debug("WeChat Error CodeInfo:", w.Errcode, w.Errmsg)
					// 	}
					// 	resp.Body.Close()
					// 	break
					// }
					resp.Body.Close()
					break
				}
				resp.Body.Close()
			} else {
				beego.Error("SMS Send msg shutdown!!! ")
			}
		}
	}()
}

func (s *SMS) sendSMSmsg(phoneNum, msg, sign string) error {
	body := fmt.Sprintf(s.USER_URL, phoneNum, msg, sign)
	s.URL += body
	select {
	case s.msgch <- s.URL:
		return nil
	default:
		beego.Error("SMS message ch full")
		return fmt.Errorf("SMS message ch full")
	}
}
