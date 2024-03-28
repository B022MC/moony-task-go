package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type IJobExpectationDAO interface {
	Get(id int64) (*model.JobExpectation, error)
	GetByUserId(userId int64) ([]*model.JobExpectation, error)
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
