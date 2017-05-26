package msg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/astaxie/beego"
	// for json get
)

var Lock sync.Mutex

type SMS struct {
	URL              string
	USER_ACCOUNT_URL string
	USER_POST_Url    string
	USER_CODE_Url    string
	msgch            chan string
	Codech           chan string
}

func Start(p *Config) *SMS {
	s := SMS{}
	s.URL = p.Url
	s.USER_ACCOUNT_URL = p.USER_ACCOUNT_URL
	s.USER_POST_Url = p.USER_POST_Url
	s.USER_CODE_Url = p.USER_IDENTI_Url
	s.msgch = make(chan string, 10240)
	s.Codech = make(chan string, 10240)
	return &s
}

func (s *SMS) RunSMSing() {
	go func() {
		for {
			url, ok := <-s.msgch
			if ok {
				postReq, err := http.NewRequest("POST", s.URL, strings.NewReader(url))
				if err != nil {
					beego.Debug("POST SMS MSG Fail", err)
					continue
				}
				postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				client := &http.Client{}
				resp, err := client.Do(postReq)
				if resp.StatusCode != http.StatusOK || err != nil {
					beego.Debug("resp.Status error: ", resp.Status, err)
					resp.Body.Close()
					continue
				}
				if resp != nil {
					buf, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						beego.Debug("Read Body error", err)
						resp.Body.Close()
						continue
					}
					str := fmt.Sprintf("%s", buf)
					code := strings.Split(str, ",")
					nCode, _ := strconv.Atoi(code[0])
					if nCode != 0 {
						beego.Debug("SMS MSG Send Fail, Error Code", nCode)
					}
					resp.Body.Close()
					continue
				}
				resp.Body.Close()
			} else {
				beego.Error("SMS MSG Send msg shutdown!!! ")
			}
		}
	}()
}

func (s *SMS) RunCodeing() {
	go func() {
		for {
			url, ok := <-s.Codech
			if ok {
				postReq, err := http.NewRequest("POST", s.URL, strings.NewReader(url))
				if err != nil {
					beego.Debug("POST SMS IDENTIFY CODE Fail", err)
					continue
				}
				postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				client := &http.Client{}
				resp, err := client.Do(postReq)
				if resp.StatusCode != http.StatusOK || err != nil {
					beego.Debug("resp.Status error: ", resp.Status, err)
					resp.Body.Close()
					continue
				}
				if resp != nil {
					buf, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						beego.Debug("Read Body error", err)
						resp.Body.Close()
						continue
					}
					str := fmt.Sprintf("%s", buf)
					code := strings.Split(str, ",")
					nCode, _ := strconv.Atoi(code[0])
					if nCode != 0 {
						beego.Debug("SMSCode Send Fail, Error Code", nCode)
					}
					resp.Body.Close()
					continue
				}
				resp.Body.Close()
			} else {
				beego.Error("SMSCode Send msg shutdown!!! ")
			}
		}
	}()
}

func (s *SMS) sendSMSmsg(phoneNum, sign, msg string) error {
	body := fmt.Sprintf(s.USER_POST_Url, phoneNum, sign, msg)
	Send_Url := s.USER_ACCOUNT_URL + body
	select {
	case s.msgch <- Send_Url:
		return nil
	default:
		beego.Error("SMS Msg message ch full")
		return fmt.Errorf("SMS Msg message ch full")
	}
}

func (s *SMS) sendSMSCode(phoneNum, sign string, code int64) error {
	body := fmt.Sprintf(s.USER_CODE_Url, phoneNum, sign, code)
	Send_Url := s.USER_ACCOUNT_URL + body
	select {
	case s.Codech <- Send_Url:
		return nil
	default:
		beego.Error("SMS Code message ch full")
		return fmt.Errorf("SMS code message ch full")
	}
}
