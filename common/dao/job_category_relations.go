package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

// JobCategoryRelationDao 定义了对 job_category_relations 表的操作方法
type JobCategoryRelationDao interface {
	AddJobCategoryRelation(relation *model.JobCategoryRelation) error
	GetRelationsByJobId(jobId int) ([]model.JobCategoryRelation, error)
	GetRelationsByCategoryId(categoryId int) ([]model.JobCategoryRelation, error)
	DeleteRelation(jobId, categoryId int) error
}
type jobCategoryRelationDaoImpl struct {
	db *gorm.DB
}

// NewJobCategoryRelationDao 创建一个新的 JobCategoryRelationDao 实例
func NewJobCategoryRelationDao(db *gorm.DB) JobCategoryRelationDao {
	return &jobCategoryRelationDaoImpl{db: db}
}

func (dao *jobCategoryRelationDaoImpl) TableName() string {
	return "job_category_relations"
}

// AddJobCategoryRelation 添加一个新的工作与类别关联
func (dao *jobCategoryRelationDaoImpl) AddJobCategoryRelation(relation *model.JobCategoryRelation) error {
	return dao.db.Table(dao.TableName()).Create(relation).Error
}

// GetRelationsByJobId 根据工作 ID 获取所有关联的类别
func (dao *jobCategoryRelationDaoImpl) GetRelationsByJobId(jobId int) ([]model.JobCategoryRelation, error) {
	var relations []model.JobCategoryRelation
	err := dao.db.Table(dao.TableName()).Where("job_id = ?", jobId).Find(&relations).Error
	return relations, err
}

// GetRelationsByCategoryId 根据类型 ID 获取所有关联的类别
func (dao *jobCategoryRelationDaoImpl) GetRelationsByCategoryId(categoryId int) ([]model.JobCategoryRelation, error) {
	var relations []model.JobCategoryRelation
	err := dao.db.Table(dao.TableName()).Where("category_id = ?", categoryId).Find(&relations).Error
	return relations, err
}

// DeleteRelation 删除一个工作与类别的关联
func (dao *jobCategoryRelationDaoImpl) DeleteRelation(jobId, categoryId int) error {
	return dao.db.Table(dao.TableName()).Where("job_id = ? AND category_id = ?", jobId, categoryId).Delete(&model.JobCategoryRelation{}).Error
}
