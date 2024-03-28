package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WxSessionResponse 定义返回的结构体
type WxSessionResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid,omitempty"`
	Errcode    int    `json:"errcode,omitempty"`
	Errmsg     string `json:"errmsg,omitempty"`
}

// Jscode2session 调用微信接口获取 session_key 和 openid
func Jscode2session(appId, secret, code string) (*WxSessionResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, secret, code)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res WxSessionResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	if res.Errcode != 0 {
		return &res, fmt.Errorf("jscode2session failed: %s", res.Errmsg)
	}

	return &res, nil
}
