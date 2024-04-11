package model

// Tag 定义了标签的结构体，对应数据库中的 tags 表
type Tag struct {
	TagId int    `json:"tag_id" form:"tagId" gorm:"primaryKey;autoIncrement"` // 标签ID
	Name  string `json:"name" form:"tagId" gorm:"type:varchar(255)"`          // 标签名称
}
