package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

type IWorkExperienceService interface {
	GetWorkExperience(id int64) (*model.WorkExperience, error)
	GetAllWorkExperiencesByUserId(userId int64) ([]*model.WorkExperience, error)
	DeleteWorkExperience(id int64) error
	CreateWorkExperience(workExperience *model.WorkExperience) (*model.WorkExperience, error)
	UpdateWorkExperience(workExperience *model.WorkExperience) (*model.WorkExperience, error)
}

type WorkExperienceService struct {
}

func NewWorkExperienceService() IWorkExperienceService {
	return &WorkExperienceService{}
}

func (s *WorkExperienceService) GetWorkExperience(id int64) (*model.WorkExperience, error) {
	return dao.NewWorkExperienceDAO(global.Db).Get(id)
}

func (s *WorkExperienceService) GetAllWorkExperiencesByUserId(userId int64) ([]*model.WorkExperience, error) {
	return dao.NewWorkExperienceDAO(global.Db).GetByUserId(userId)
}

func (s *WorkExperienceService) DeleteWorkExperience(id int64) error {
	return dao.NewWorkExperienceDAO(global.Db).Delete(id)
}

func (s *WorkExperienceService) CreateWorkExperience(workExperience *model.WorkExperience) (*model.WorkExperience, error) {
	return dao.NewWorkExperienceDAO(global.Db).Create(workExperience)
}

func (s *WorkExperienceService) UpdateWorkExperience(workExperience *model.WorkExperience) (*model.WorkExperience, error) {
	return dao.NewWorkExperienceDAO(global.Db).Update(workExperience)
}
