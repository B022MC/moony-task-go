package model

type Jobs struct {
	JobId       int64   `json:"job_id" form:"jobId"`            // 兼职信息ID
	UserId      int64   `json:"user_id" form:"userId"`          // 发布用户ID
	Title       string  `json:"title" form:"title"`             // 兼职标题
	Description string  `json:"description" form:"description"` // 兼职描述
	Location    string  `json:"location" form:"location"`       // 工作地点
	Salary      string  `json:"salary" form:"salary"`           // 薪资
	Status      int     `json:"status" form:"status"`           // 状态：1=开放 2=关闭
	Lat         float64 `json:"lat" form:"lat"`                 // 纬度
	Lng         float64 `json:"lng" form:"lng"`                 // 经度
	CreateTime  int64   `json:"create_time" form:"createTime"`  // 创建时间
	UpdateTime  int64   `json:"update_time" form:"updateTime"`  // 更新时间
}
