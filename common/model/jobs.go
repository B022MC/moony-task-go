package model

type Jobs struct {
	JobId             int64   `json:"job_id" form:"jobId"`                         // 兼职信息ID
	UserId            int64   `json:"user_id" form:"userId"`                       // 发布用户ID
	Title             string  `json:"title" form:"title"`                          // 兼职标题
	Description       string  `json:"description" form:"description"`              // 兼职描述
	Location          string  `json:"location" form:"location"`                    // 工作地点
	Salary            string  `json:"salary" form:"salary"`                        // 薪资
	CategoryId        int     `json:"category_id" form:"categoryId"`               // 类别ID
	Status            int     `json:"status" form:"status"`                        // 状态：1=开放 2=关闭
	GenderRequirement string  `json:"gender_requirement" form:"genderRequirement"` // 性别要求
	WorkPeriod        string  `json:"work_period" form:"workPeriod"`               // 工作周期
	Lat               float64 `json:"lat" form:"lat"`                              // 纬度
	Lng               float64 `json:"lng" form:"lng"`                              // 经度
	CreateTime        int64   `json:"create_time" form:"createTime"`               // 创建时间
	UpdateTime        int64   `json:"update_time" form:"updateTime"`               // 更新时间
}

// JobsAddRequest 定义了客户端提交兼职信息请求的结构体
type JobsAddRequest struct {
	UserId            int64   `json:"user_id" form:"userId" binding:"required"`    // 发布用户ID
	Title             string  `json:"title" form:"title" binding:"required"`       // 兼职标题
	Description       string  `json:"description" form:"description"`              // 兼职描述，可选
	Location          string  `json:"location" form:"location"`                    // 工作地点，可选
	Salary            string  `json:"salary" form:"salary"`                        // 薪资，可选
	CategoryId        int     `json:"category_id" form:"categoryId"`               // 类别ID
	Status            int     `json:"status" form:"status"`                        // 状态：1=开放 2=关闭
	GenderRequirement string  `json:"gender_requirement" form:"genderRequirement"` // 性别要求，可选
	WorkPeriod        string  `json:"work_period" form:"workPeriod"`               // 工作周期，可选
	Lat               float64 `json:"lat" form:"lat"`                              // 纬度，可选
	Lng               float64 `json:"lng" form:"lng"`                              // 经度，可选
}
