package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

// IndustryTypeDao 定义了对 industry_types 表的操作方法
type IndustryTypeDao interface {
	AddIndustryType(industry *model.IndustryType) error
	GetIndustryTypeById(id int) (*model.IndustryType, error)
	UpdateIndustryType(industry *model.IndustryType) error
	DeleteIndustryType(id int) error
	GetAllIndustryTypes() ([]model.IndustryType, error)
	GetAllIndustriesWithCategories() ([]model.IndustryWithCategoriesDropdown, error)
}
type industryTypeDaoImpl struct {
	db *gorm.DB
}

// NewIndustryTypeDao 创建一个新的 IndustryTypeDao 实例
func NewIndustryTypeDao(db *gorm.DB) IndustryTypeDao {
	return &industryTypeDaoImpl{db: db}
}

func (dao *industryTypeDaoImpl) TableName() string {
	return "industry_types"
}

// AddIndustryType 添加一个新的行业类型
func (dao *industryTypeDaoImpl) AddIndustryType(industry *model.IndustryType) error {
	return dao.db.Create(industry).Error
}

// GetIndustryTypeById 根据 ID 获取行业类型
func (dao *industryTypeDaoImpl) GetIndustryTypeById(id int) (*model.IndustryType, error) {
	var industry model.IndustryType
	result := dao.db.Table(dao.TableName()).First(&industry, id)
	return &industry, result.Error
}

// UpdateIndustryType 更新行业类型信息
func (dao *industryTypeDaoImpl) UpdateIndustryType(industry *model.IndustryType) error {
	return dao.db.Table(dao.TableName()).Save(industry).Error
}

// DeleteIndustryType 删除一个行业类型
func (dao *industryTypeDaoImpl) DeleteIndustryType(id int) error {
	return dao.db.Table(dao.TableName()).Delete(&model.IndustryType{}, id).Error
}

// GetAllIndustryTypes 获取所有行业类型
func (dao *industryTypeDaoImpl) GetAllIndustryTypes() ([]model.IndustryType, error) {
	var industries []model.IndustryType
	err := dao.db.Table(dao.TableName()).Find(&industries).Error
	return industries, err
}

func (dao *industryTypeDaoImpl) GetAllIndustriesWithCategories() ([]model.IndustryWithCategoriesDropdown, error) {
	var industries []model.IndustryType
	err := dao.db.Find(&industries).Error
	if err != nil {
		return nil, err
	}

	var result []model.IndustryWithCategoriesDropdown
	for _, industry := range industries {
		var relations []model.CategoryIndustryRelation
		err = dao.db.Table("category_industry_relations").Where("industry_id = ?", industry.IndustryId).Find(&relations).Error
		if err != nil {
			return nil, err
		}

		var children []model.Child
		for _, rel := range relations {
			var category model.JobCategory
			err = dao.db.Where("category_id = ?", rel.CategoryId).Find(&category).Error
			if err != nil {
				return nil, err
			}
			child := model.Child{
				CategoryId: category.CategoryId, // 添加类别ID
				Label:      category.Name,
				Check:      false, // 默认为 false，可以根据具体需求调整
			}
			children = append(children, child)
		}

		dropdown := model.IndustryWithCategoriesDropdown{
			Title:    industry.Name,
			Children: children,
		}
		result = append(result, dropdown)
	}

	return result, nil
}
