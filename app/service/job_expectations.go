package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

type IJobExpectationService interface {
	GetJobExpectation(id int64) (*model.JobExpectation, error)
	GetAllJobExpectationsByUserId(userId int64) ([]*model.JobExpectation, error)
	DeleteJobExpectation(id int64) error
	CreateJobExpectation(jobExpectation *model.JobExpectation) (*model.JobExpectation, error)
	UpdateJobExpectation(jobExpectation *model.JobExpectation) (*model.JobExpectation, error)
}

type JobExpectationService struct{}

func NewJobExpectationService() IJobExpectationService {
	return &JobExpectationService{}
}

func (s *JobExpectationService) GetJobExpectation(id int64) (*model.JobExpectation, error) {
	return dao.NewJobExpectationDAO(global.Db).Get(id)
}

func (s *JobExpectationService) GetAllJobExpectationsByUserId(userId int64) ([]*model.JobExpectation, error) {
	return dao.NewJobExpectationDAO(global.Db).GetByUserId(userId)
}

func (s *JobExpectationService) DeleteJobExpectation(id int64) error {
	return dao.NewJobExpectationDAO(global.Db).Delete(id)
}

func (s *JobExpectationService) CreateJobExpectation(jobExpectation *model.JobExpectation) (*model.JobExpectation, error) {
	return dao.NewJobExpectationDAO(global.Db).Create(jobExpectation)
}

func (s *JobExpectationService) UpdateJobExpectation(jobExpectation *model.JobExpectation) (*model.JobExpectation, error) {
	return dao.NewJobExpectationDAO(global.Db).Update(jobExpectation)
}
