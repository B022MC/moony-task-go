package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"net/http"
)

const (
	code2AccessTokenUrl     string = "https://api.weixin.qq.com/sns/oauth2/access_token"
	accessToken2UserInfoUrl string = "https://api.weixin.qq.com/sns/userinfo"
)

type OpenSDK struct {
	AppId  string
	Secret string
}

type UserInfo struct {
	Openid      string
	AccessToken string
	Unionid     string
	Nickname    string
	HeadUrl     string
}

func NewOpenSdk(appid, secret string) *OpenSDK {
	return &OpenSDK{
		AppId:  appid,
		Secret: secret,
	}
}

func (o *OpenSDK) GetUserInfo(code string) (*UserInfo, error) {
	url := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		code2AccessTokenUrl, o.AppId, o.Secret, code)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]interface{})
	err = json.Unmarshal(body, &mp)
	if err != nil {
		return nil, err
	}

	errcode := cast.ToInt64(mp["errcode"])
	if errcode != 0 {
		return nil, fmt.Errorf("%d:%s", errcode, cast.ToString(mp["errmsg"]))
	}

	var user UserInfo
	user.Unionid = cast.ToString(mp["unionid"])
	user.AccessToken = cast.ToString(mp["access_token"])
	user.Openid = cast.ToString(mp["openid"])

	err = o.getUserInfo(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (o *OpenSDK) getUserInfo(user *UserInfo) error {
	url := fmt.Sprintf("%s?access_token=%s&openid=%s",
		accessToken2UserInfoUrl, user.AccessToken, user.Openid)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	mp := make(map[string]interface{})
	err = json.Unmarshal(body, &mp)
	if err != nil {
		return err
	}
	errcode := cast.ToInt64(mp["errcode"])
	if errcode != 0 {
		return fmt.Errorf("%d:%s", errcode, cast.ToString(mp["errmsg"]))
	}

	user.Nickname = cast.ToString(mp["nickname"])
	user.HeadUrl = cast.ToString(mp["headimgurl"])
	return nil
}
