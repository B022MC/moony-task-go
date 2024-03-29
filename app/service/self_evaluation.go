package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

type ISelfEvaluationService interface {
	GetSelfEvaluation(id int64) (*model.SelfEvaluation, error)
	GetAllSelfEvaluationsByUserId(userId int64) ([]*model.SelfEvaluation, error)
	DeleteSelfEvaluation(id int64) error
	CreateSelfEvaluation(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error)
	UpdateSelfEvaluation(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error)
}

type SelfEvaluationService struct{}

func NewSelfEvaluationService() ISelfEvaluationService {
	return &SelfEvaluationService{}
}

func (s *SelfEvaluationService) GetSelfEvaluation(id int64) (*model.SelfEvaluation, error) {
	return dao.NewSelfEvaluationDAO(global.Db).Get(id)
}

func (s *SelfEvaluationService) GetAllSelfEvaluationsByUserId(userId int64) ([]*model.SelfEvaluation, error) {
	return dao.NewSelfEvaluationDAO(global.Db).GetByUserId(userId)
}

func (s *SelfEvaluationService) DeleteSelfEvaluation(id int64) error {
	return dao.NewSelfEvaluationDAO(global.Db).Delete(id)
}

func (s *SelfEvaluationService) CreateSelfEvaluation(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error) {
	return dao.NewSelfEvaluationDAO(global.Db).Create(selfEvaluation)
}

func (s *SelfEvaluationService) UpdateSelfEvaluation(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error) {
	return dao.NewSelfEvaluationDAO(global.Db).Update(selfEvaluation)
}
