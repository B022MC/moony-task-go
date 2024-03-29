package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type IJobExpectationDAO interface {
	Get(id int64) (*model.JobExpectation, error)
	GetByUserId(userId int64) ([]*model.JobExpectation, error)
	Delete(id int64) error
	Create(jobExpectation *model.JobExpectation) (*model.JobExpectation, error)
	Update(jobExpectation *model.JobExpectation) (*model.JobExpectation, error)
}

type JobExpectationDAO struct {
	DB *gorm.DB
}

func NewJobExpectationDAO(db *gorm.DB) IJobExpectationDAO {
	return &JobExpectationDAO{DB: db}
}

func (dao *JobExpectationDAO) Get(id int64) (*model.JobExpectation, error) {
	var jobExpectation model.JobExpectation
	result := dao.DB.Where("id = ?", id).First(&jobExpectation)
	return &jobExpectation, result.Error
}

func (dao *JobExpectationDAO) GetByUserId(userId int64) ([]*model.JobExpectation, error) {
	var jobExpectations []*model.JobExpectation
	result := dao.DB.Where("user_id = ?", userId).Find(&jobExpectations)
	return jobExpectations, result.Error
}
func (dao *JobExpectationDAO) Delete(id int64) error {
	result := dao.DB.Delete(&model.JobExpectation{}, id)
	return result.Error
}

func (dao *JobExpectationDAO) Create(jobExpectation *model.JobExpectation) (*model.JobExpectation, error) {
	result := dao.DB.Create(jobExpectation)
	return jobExpectation, result.Error
}

func (dao *JobExpectationDAO) Update(jobExpectation *model.JobExpectation) (*model.JobExpectation, error) {
	result := dao.DB.Save(jobExpectation)
	return jobExpectation, result.Error
}
