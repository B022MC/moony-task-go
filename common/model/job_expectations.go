package model

type JobExpectation struct {
	Id                   int64  `json:"id" form:"id"`                                        // 工作期望ID
	UserId               int64  `json:"user_id" form:"userId"`                               // 用户ID
	WorkType             string `json:"work_type" form:"workType"`                           // 期望工作类型
	AvailableDate        string `json:"available_date" form:"availableDate"`                 // 期望工作日期
	AvailableTimePerWeek string `json:"available_time_per_week" form:"availableTimePerWeek"` // 每周可上班时间
	FullTimeAvailable    bool   `json:"full_time_available" form:"fullTimeAvailable"`        // 是否支持全职上班
	CreateTime           int64  `json:"create_time" form:"createTime"`                       // 创建时间
	UpdateTime           int64  `json:"update_time" form:"updateTime"`                       // 更新时间
}

// JobExpectationFormatted 定义了工作期望的格式化输出结构
type JobExpectationFormatted struct {
	Id                   int64  `json:"id"`
	UserId               int64  `json:"user_id"`
	WorkType             string `json:"work_type"`
	AvailableDate        string `json:"available_date"`
	AvailableTimePerWeek string `json:"available_time_per_week"`
	FullTimeAvailable    string `json:"full_time_available"`
}

// Format 方法将 JobExpectation 实例格式化为 JobExpectationFormatted
func (je *JobExpectation) Format() JobExpectationFormatted {
	fullTimeAvailableStr := "否"
	if je.FullTimeAvailable {
		fullTimeAvailableStr = "是"
	}

	return JobExpectationFormatted{
		Id:                   je.Id,
		UserId:               je.UserId,
		WorkType:             je.WorkType,
		AvailableDate:        je.AvailableDate,
		AvailableTimePerWeek: je.AvailableTimePerWeek,
		FullTimeAvailable:    fullTimeAvailableStr,
	}
}
