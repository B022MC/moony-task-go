package model

// JobTag 定义了工作和标签之间关系的结构体，对应数据库中的 job_tags 表
type JobTag struct {
	JobId int64 `json:"job_id" gorm:"primaryKey;index"` // 兼职信息ID
	TagId int   `json:"tag_id" gorm:"primaryKey;index"` // 标签ID
}
