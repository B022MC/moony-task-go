package model

import (
	"fmt"
	"time"
)

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

// WorkExperienceFormatted 定义了工作经历的格式化输出结构
type WorkExperienceFormatted struct {
	Id             int64  `json:"id"`
	UserId         int64  `json:"user_id"`
	TimePeriod     string `json:"time_period"` // 易读的工作期间
	CompanyName    string `json:"company_name"`
	JobDescription string `json:"job_description"`
	CreateTime     string `json:"create_time"` // 格式化时间
	UpdateTime     string `json:"update_time"` // 格式化时间
}

func (we *WorkExperience) Format() WorkExperienceFormatted {
	// 假设 StartTime 和 EndTime 已是易读格式或进行相应的格式化
	// 对于创建时间和更新时间，进行格式化为更易读的字符串格式
	createTimeStr := time.Unix(we.CreateTime, 0).Format("2006-01-02 15:04:05")
	updateTimeStr := time.Unix(we.UpdateTime, 0).Format("2006-01-02 15:04:05")

	return WorkExperienceFormatted{
		Id:             we.Id,
		UserId:         we.UserId,
		TimePeriod:     fmt.Sprintf("%s - %s", we.StartTime, we.EndTime),
		CompanyName:    we.CompanyName,
		JobDescription: we.JobDescription,
		CreateTime:     createTimeStr,
		UpdateTime:     updateTimeStr,
	}
}
