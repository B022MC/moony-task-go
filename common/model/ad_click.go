package model

import "moony-task-go/utils"

type ClientInfo struct {
	AppId   int64  `form:"appId" json:"appId"`
	Channel string `form:"channel" json:"channel"`
	Version string `form:"version" json:"version"`
	Os      string `form:"os" json:"os"`
	Ip      string `form:"ip" json:"ip"`
	Ua      string `form:"ua" json:"ua"`
	Model   string `form:"model" json:"model"`
	Idfa    string `form:"idfa" json:"idfa"`
	Oaid    string `form:"oaid" json:"oaid"`
	Imei    string `form:"imei" json:"imei"`
	Extra   string `form:"extra" json:"-"`
	Active  bool   `form:"active" json:"active"`
}

type AdBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AdId   int64  `json:"adId"`
		Source string `json:"source"`
		Extra  struct {
			MatchType  string `json:"matchType"`
			AdPlanId   string `json:"adPlanId"`
			AdPlanName string `json:"adPlanName"`
			MaterialId string `json:"materialId"`
		} `json:"extra"`
	} `json:"data"`
}

func (c *ClientInfo) Repair() {
	if c.Oaid == "00000000-0000-0000-0000-000000000000" || c.Oaid == "undefined" {
		c.Oaid = ""
	}
	if c.Model == "" && c.Ua != "" {
		c.Model = utils.ParseModelByUa(c.Ua)
	}
	if c.Os == "" {
		c.Os = "android"
	}
}
