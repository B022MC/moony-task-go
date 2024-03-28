package model

type Education struct {
	Id         int64  `json:"id" form:"id"`                  // 教育经历ID
	UserId     int64  `json:"user_id" form:"userId"`         // 用户ID
	StartTime  string `json:"start_time" form:"startTime"`   // 开始时间
	EndTime    string `json:"end_time" form:"endTime"`       // 结束时间
	SchoolName string `json:"school_name" form:"schoolName"` // 学校名称
	Degree     string `json:"degree" form:"degree"`          // 学位
	Major      string `json:"major" form:"major"`            // 专业
	CreateTime int64  `json:"create_time" form:"createTime"` // 创建时间
	UpdateTime int64  `json:"update_time" form:"updateTime"` // 更新时间
}
