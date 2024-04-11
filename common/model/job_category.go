package model

// JobCategory 定义了工作类别的结构体，对应数据库中的 job_categories 表
type JobCategory struct {
	CategoryId int    `json:"category_id" gorm:"primaryKey;autoIncrement"` // 类别ID
	Name       string `json:"name" gorm:"type:varchar(255)"`               // 类别名称
}
