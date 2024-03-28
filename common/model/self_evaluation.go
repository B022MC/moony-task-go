package model

type SelfEvaluation struct {
	Id         int64  `json:"id" form:"id"`                  // 自我评价ID
	UserId     int64  `json:"user_id" form:"userId"`         // 用户ID
	Content    string `json:"content" form:"content"`        // 自我评价内容
	CreateTime int64  `json:"create_time" form:"createTime"` // 创建时间
	UpdateTime int64  `json:"update_time" form:"updateTime"` // 更新时间
}
