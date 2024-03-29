package model

import "time"

type SelfEvaluation struct {
	Id         int64  `json:"id" form:"id"`                  // 自我评价ID
	UserId     int64  `json:"user_id" form:"userId"`         // 用户ID
	Content    string `json:"content" form:"content"`        // 自我评价内容
	CreateTime int64  `json:"create_time" form:"createTime"` // 创建时间
	UpdateTime int64  `json:"update_time" form:"updateTime"` // 更新时间
}

// SelfEvaluationFormatted 定义了自我评价的格式化输出结构
type SelfEvaluationFormatted struct {
	Id         int64  `json:"id"`
	UserId     int64  `json:"user_id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"` // 格式化时间
	UpdateTime string `json:"update_time"` // 格式化时间
}

func (se *SelfEvaluation) Format() SelfEvaluationFormatted {
	// 假设创建时间和更新时间是Unix时间戳，我们将其转换为更易读的格式
	createTimeStr := time.Unix(se.CreateTime, 0).Format("2006-01-02 15:04:05")
	updateTimeStr := time.Unix(se.UpdateTime, 0).Format("2006-01-02 15:04:05")

	return SelfEvaluationFormatted{
		Id:         se.Id,
		UserId:     se.UserId,
		Content:    se.Content,
		CreateTime: createTimeStr,
		UpdateTime: updateTimeStr,
	}
}
