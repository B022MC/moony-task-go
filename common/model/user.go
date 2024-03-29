package model

var (
	UserStatusEnable  = 1
	UserStatusDisable = 2

	LoginTypeWeixin = "weixin"    //微信登录
	LoginTypePhone  = "phone"     //手机号登录
	LoginUniVerify  = "univerify" //uniapp一键手机号登陆

)

type User struct {
	Id         int64  `json:"id" form:"id"`                  // 用户id
	AppId      int64  `json:"app_id" form:"appId"`           // 应用ID
	HashId     uint64 `json:"hash_id" form:"hashId"`         // 用户hash_id
	HashNum    int    `json:"hash_num" form:"hashNum"`       // 哈希计算值
	Status     int    `json:"status" form:"status"`          // 状态：1=正常 2=禁用
	NickName   string `json:"nick_name" form:"nickName"`     // 昵称
	Name       string `json:"name" form:"name"`              // 名字
	Config     string `json:"config" form:"config"`          // 用户个性化配置
	Avatar     string `json:"avatar" form:"avatar"`          // 头像
	Unionid    string `json:"unionid" form:"unionid"`        // 微信统一ID
	Phone      string `json:"phone" form:"phone"`            // 手机号
	Age        int    `json:"age" form:"age"`                // 年龄
	Sex        int    `json:"sex" form:"sex"`                // 性别：1=男，2=女，3=未知
	CreateTime int64  `json:"create_time" form:"createTime"` // 创建时间
	ActiveTime int64  `json:"active_time" form:"activeTime"` // 活跃时间
	UpdateTime int64  `json:"update_time" form:"updateTime"` // 更新时间
	TokenMd5   string `json:"token_md5" form:"tokenMd5"`     // 用户token md5值
}

type UserRes struct {
	Id      int64  `json:"id"`
	AppId   int64  `json:"appId"`
	HashId  uint64 `json:"hashId"`
	HashNum int    `json:"hashNum"`
	Role    int    `json:"role"`
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Config  string `json:"config"`
	Avatar  string `json:"avatar"`
	Phone   string `json:"phone"`
	//Business   string `json:"business"`
	//Platform   string `json:"platform"`
	//Version    string `json:"version"`
	//Channel    string `json:"channel"`
	//Brand      string `json:"brand"`
	//Model      string `json:"model"`
	CreateTime string `json:"createTime"`
	ActiveTime string `json:"activeTime"`
	UpdateTime string `json:"updateTime"`
}

type UserConfigReq struct {
	Oaid string `form:"oaid"`
	Idfa string `form:"idfa"`
	Imei string `form:"imei"`
	Ua   string `form:"ua"`
}

type UserConfigRsp struct {
	Config map[string]interface{} `json:"config"`
	UserId string                 `json:"userId"`
	HashId string                 `json:"hashId"`
	Guest  bool                   `json:"guest"`
	Token  string                 `json:"token"`
	Debug  interface{}            `json:"debug"`
}

type LoginUserReq struct {
	LoginType string `form:"loginType" json:"loginType"`
	Weixin    struct {
		Code     string `form:"code" json:"code"`
		UserInfo struct {
			NickName  string `json:"nickName"`
			AvatarUrl string `json:"avatarUrl"`
		} `json:"userInfo"`
	} `form:"weixin" json:"weixin"`
}

func (u *User) IsGuest() bool {
	if u.Unionid == "" && u.Phone == "" {
		return true
	}
	return false
}
