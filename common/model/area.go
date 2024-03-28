package model

type Area struct {
	Id         int    `json:"id" form:"id"`                   // ID
	Pid        int    `json:"pid" form:"pid"`                 // 父id
	ShortName  string `json:"shortname" form:"shortname"`     // 简称
	Name       string `json:"name" form:"name"`               // 名称
	MergerName string `json:"merger_name" form:"merger_name"` // 全称
	Level      int    `json:"level" form:"level"`             // 层级 0 1 2 省市区县
	Pinyin     string `json:"pinyin" form:"pinyin"`           // 拼音
	Code       string `json:"code" form:"code"`               // 长途区号
	ZipCode    string `json:"zip_code" form:"zip_code"`       // 邮编
	First      string `json:"first" form:"first"`             // 首字母
	Lng        string `json:"lng" form:"lng"`                 // 经度
	Lat        string `json:"lat" form:"lat"`                 // 纬度
}
