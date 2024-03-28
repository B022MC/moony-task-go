package model

var (
	DefaultAvatar = "https://cdn.yic3.cn/recruit/%E5%B7%A5%E4%BA%BA%E9%BB%98%E8%AE%A4%E5%A4%B4%E5%83%8F.png"
)

type Work struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Pid         string  `json:"pid"`
	Level       string  `json:"level"`
	Letter      string  `json:"letter"`
	SpecialType string  `json:"specialType"`
	ExtType     string  `json:"extType"`
	ShortAlias  string  `json:"shortAlias"`
	ShowOrder   string  `json:"showOrder"`
	Children    []*Work `json:"children"`
}
