package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

type IEducationService interface {
	GetEducation(id int64) (*model.Education, error)
	GetAllEducationsByUserId(userId int64) ([]*model.Education, error)
	DeleteEducation(id int64) error
	CreateEducation(education *model.Education) (*model.Education, error)
	UpdateEducation(education *model.Education) (*model.Education, error)
}

type EducationService struct{}

func NewEducationService() IEducationService {
	return &EducationService{}
}

func (s *EducationService) GetEducation(id int64) (*model.Education, error) {
	return dao.NewEducationDAO(global.Db).Get(id)
}

func (s *EducationService) GetAllEducationsByUserId(userId int64) ([]*model.Education, error) {
	return dao.NewEducationDAO(global.Db).GetByUserId(userId)
}

func (s *EducationService) DeleteEducation(id int64) error {
	return dao.NewEducationDAO(global.Db).Delete(id)
}

func (s *EducationService) CreateEducation(education *model.Education) (*model.Education, error) {
	return dao.NewEducationDAO(global.Db).Create(education)
}

func (s *EducationService) UpdateEducation(education *model.Education) (*model.Education, error) {
	return dao.NewEducationDAO(global.Db).Update(education)
}
