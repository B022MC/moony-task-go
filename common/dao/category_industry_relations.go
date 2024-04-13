package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

// CategoryIndustryRelationDao 定义了对 category_industry_relations 表的操作方法
type CategoryIndustryRelationDao interface {
	AddRelation(relation *model.CategoryIndustryRelation) error
	GetRelationsByCategoryId(categoryId int) ([]model.CategoryIndustryRelation, error)
	DeleteRelation(categoryId, industryId int) error
}

type categoryIndustryRelationDaoImpl struct {
	db *gorm.DB
}

// NewCategoryIndustryRelationDao 创建一个新的 CategoryIndustryRelationDao 实例
func NewCategoryIndustryRelationDao(db *gorm.DB) CategoryIndustryRelationDao {
	return &categoryIndustryRelationDaoImpl{db: db}
}
func (dao *categoryIndustryRelationDaoImpl) TableName() string {
	return "category_industry_relations"
}

// AddRelation 添加一个新的类别与行业的关联
func (dao *categoryIndustryRelationDaoImpl) AddRelation(relation *model.CategoryIndustryRelation) error {
	return dao.db.Table(dao.TableName()).Create(relation).Error
}

// GetRelationsByCategoryId 根据类别 ID 获取所有关联的行业
func (dao *categoryIndustryRelationDaoImpl) GetRelationsByCategoryId(categoryId int) ([]model.CategoryIndustryRelation, error) {
	var relations []model.CategoryIndustryRelation
	err := dao.db.Table(dao.TableName()).Where("category_id = ?", categoryId).Find(&relations).Error
	return relations, err
}

// DeleteRelation 删除一个类别与行业的关联
func (dao *categoryIndustryRelationDaoImpl) DeleteRelation(categoryId, industryId int) error {
	return dao.db.Table(dao.TableName()).Where("category_id = ? AND industry_id = ?", categoryId, industryId).Delete(&model.CategoryIndustryRelation{}).Error
}
