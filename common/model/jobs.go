package model

type Jobs struct {
	JobId             int64   `json:"job_id" form:"jobId"`                         // 兼职信息ID
	UserId            int64   `json:"user_id" form:"userId"`                       // 发布用户ID
	Title             string  `json:"title" form:"title"`                          // 兼职标题
	Description       string  `json:"description" form:"description"`              // 兼职描述
	Location          string  `json:"location" form:"location"`                    // 工作地点描述
	CityId            int     `json:"city_id" form:"cityId"`                       // 城市id
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

// JobCategoryRelation 定义了工作与类别关系的结构体，对应数据库中的 job_category_relations 表
type JobCategoryRelation struct {
	JobId      int `json:"job_id" gorm:"primaryKey"`      // 工作ID
	CategoryId int `json:"category_id" gorm:"primaryKey"` // 类别ID
}

// JobFilterRequest 定义了客户端提交兼职信息请求的结构体
type JobFilterRequest struct {
	CategoryIds string            `json:"category_ids"`
	Area        AreaFilter        `json:"area"`
	Sort        Sort              `json:"sort"`
	Filters     []FilterCriterion `json:"filters"`
}
type AreaFilter struct {
	CityId string `json:"city_id"`
	Level  int64  `json:"level"`
}

type FilterCriterion struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type Sort struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	Desc string  `json:"desc"`
}

// JobRsp represents the structure of a job listing.
type JobRsp struct {
	Id          int64  `json:"id"`          // 职位ID
	Title       string `json:"title"`       // 职位标题
	Description string `json:"description"` // 职位描述
	Salary      string `json:"salary"`      // 薪资
	Location    string `json:"location"`    // 工作地点
}

// JobsRsp is a slice of JobRsp, used for a collection of jobs.
type JobsRsp []JobRsp
