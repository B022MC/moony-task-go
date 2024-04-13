package model

// IndustryType 定义了行业类型的结构体，对应数据库中的 industry_types 表
type IndustryType struct {
	IndustryId int    `json:"industry_id" gorm:"primaryKey;autoIncrement"` // 行业ID
	Name       string `json:"name" gorm:"type:varchar(255)"`               // 行业名称
}

// JobCategory 定义了工作类别的结构体，对应数据库中的 job_categories 表
type JobCategory struct {
	CategoryId int    `json:"category_id" gorm:"primaryKey;autoIncrement"` // 类别ID
	Name       string `json:"name" gorm:"type:varchar(255)"`               // 类别名称
}

// CategoryIndustryRelation defines the structure for the relationship between job categories and industry types,
// corresponding to the `category_industry_relations` table in the database.
type CategoryIndustryRelation struct {
	CategoryId int `json:"category_id" gorm:"primaryKey;foreignKey:CategoryId;association_foreignkey:CategoryID"` // JobRsp Category ID
	IndustryId int `json:"industry_id" gorm:"primaryKey;foreignKey:IndustryId;association_foreignkey:IndustryID"` // Industry CategoryIds ID
}

// IndustryWithCategoriesDropdown 用于表示下拉菜单中的行业及其类别
type IndustryWithCategoriesDropdown struct {
	Title    string  `json:"title"`
	Children []Child `json:"children"`
}

// Child 用于表示下拉菜单中的具体工作类别
type Child struct {
	CategoryId int    `json:"category_id"` // 类别ID
	Label      string `json:"label"`
	Check      bool   `json:"check"`
}
