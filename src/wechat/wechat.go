package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	// for json get
)

type Wechat struct {
	appID                string
	appSecret            string
	accessTokenFetchUrl  string
	customServicePostUrl string
}

type AccessTokenRsp struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type AccessTokenErrRsp struct {
	Errcode float64 `json:"errcode"`
	Errmsg  string  `json:"errmsg"`
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

func getParam() *Wechat {
	var config Wechat
	config.appID = beego.AppConfig.String("APPID")
	config.appSecret = beego.AppConfig.String("APPSECRET")
	config.accessTokenFetchUrl = beego.AppConfig.String("TOKEN_URL")
	config.customServicePostUrl = beego.AppConfig.String("CUSOMSER_POST_URL")
	return &config
}

func GetAccessToken() (string, float64, error) {

	pConfig := getParam()

	requestLine := fmt.Sprintf(pConfig.accessTokenFetchUrl, pConfig.appID, pConfig.appSecret)
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
		Rsp := AccessTokenRsp{}
		err = json.Unmarshal(body, &Rsp)
		if err != nil {
			return "", 0.0, err
		}
		return Rsp.AccessToken, Rsp.ExpiresIn, nil
	} else {
		beego.Debug("Request error!!!!!!!")
		Rsp := AccessTokenErrRsp{}
		err = json.Unmarshal(body, &Rsp)
		if err != nil {
			return "", 0.0, err
		}
		return "", 0.0, fmt.Errorf("%s", Rsp.Errmsg)
	}
}

func SendCustomTxTMsg(Token, openId, msg string) error {
	pConfig := getParam()

	TxtMsg := &CustomMsg{
		ToUser:  openId,
		MsgType: "text",
		Text:    TextMsgContent{Content: msg},
	}

	body, err := json.MarshalIndent(TxtMsg, " ", "  ")
	if err != nil {
		return err
	}

	TextUrl := fmt.Sprintf(pConfig.customServicePostUrl, Token)
	beego.Debug("http POST Text:", TextUrl)

	postReq, err := http.NewRequest("POST", TextUrl, bytes.NewReader(body))
	if err != nil {
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
