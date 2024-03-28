package model

type WorkExperience struct {
	Id             int64  `json:"id" form:"id"`                          // 工作经历ID
	UserId         int64  `json:"user_id" form:"userId"`                 // 用户ID
	StartTime      string `json:"start_time" form:"startTime"`           // 开始时间
	EndTime        string `json:"end_time" form:"endTime"`               // 结束时间
	CompanyName    string `json:"company_name" form:"companyName"`       // 公司名称
	JobDescription string `json:"job_description" form:"jobDescription"` // 工作描述
	CreateTime     int64  `json:"create_time" form:"createTime"`         // 创建时间
	UpdateTime     int64  `json:"update_time" form:"updateTime"`         // 更新时间
}
