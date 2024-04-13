package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

// IndustryTypeService 定义了行业类型的服务操作
type IndustryTypeService interface {
	AddIndustryType(industry *model.IndustryType) error
	GetIndustryTypeById(id int) (*model.IndustryType, error)
	UpdateIndustryType(industry *model.IndustryType) error
	DeleteIndustryType(id int) error
	GetAllIndustryTypes() ([]model.IndustryType, error)
	GetAllIndustriesWithCategories() ([]model.IndustryWithCategoriesDropdown, error)
}

type industryTypeServiceImpl struct {
}

// NewIndustryTypeService 创建一个新的 IndustryTypeService 实例
func NewIndustryTypeService() IndustryTypeService {
	return &industryTypeServiceImpl{}
}

// AddIndustryType 添加一个新的行业类型
func (s *industryTypeServiceImpl) AddIndustryType(industry *model.IndustryType) error {
	return dao.NewIndustryTypeDao(global.Db).AddIndustryType(industry)
}

// GetIndustryTypeById 根据 ID 获取行业类型
func (s *industryTypeServiceImpl) GetIndustryTypeById(id int) (*model.IndustryType, error) {
	return dao.NewIndustryTypeDao(global.Db).GetIndustryTypeById(id)
}

// UpdateIndustryType 更新行业类型信息
func (s *industryTypeServiceImpl) UpdateIndustryType(industry *model.IndustryType) error {
	return dao.NewIndustryTypeDao(global.Db).UpdateIndustryType(industry)
}

// DeleteIndustryType 删除一个行业类型
func (s *industryTypeServiceImpl) DeleteIndustryType(id int) error {
	return dao.NewIndustryTypeDao(global.Db).DeleteIndustryType(id)
}

// GetAllIndustryTypes 获取所有行业类型
func (s *industryTypeServiceImpl) GetAllIndustryTypes() ([]model.IndustryType, error) {
	return dao.NewIndustryTypeDao(global.Db).GetAllIndustryTypes()
}

// GetAllIndustriesWithCategories 获取所有行业类型及其关联的工作类别
func (s *industryTypeServiceImpl) GetAllIndustriesWithCategories() ([]model.IndustryWithCategoriesDropdown, error) {
	return dao.NewIndustryTypeDao(global.Db).GetAllIndustriesWithCategories()
}
