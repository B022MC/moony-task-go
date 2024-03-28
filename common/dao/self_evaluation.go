package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type ISelfEvaluationDAO interface {
	Get(id int64) (*model.SelfEvaluation, error)
	GetByUserId(userId int64) ([]*model.SelfEvaluation, error)
}

type SelfEvaluationDAO struct {
	DB *gorm.DB
}

func NewSelfEvaluationDAO(db *gorm.DB) ISelfEvaluationDAO {
	return &SelfEvaluationDAO{DB: db}
}

func (dao *SelfEvaluationDAO) Get(id int64) (*model.SelfEvaluation, error) {
	var selfEvaluation model.SelfEvaluation
	result := dao.DB.Where("id = ?", id).First(&selfEvaluation)
	return &selfEvaluation, result.Error
}

func (dao *SelfEvaluationDAO) GetByUserId(userId int64) ([]*model.SelfEvaluation, error) {
	var selfEvaluations []*model.SelfEvaluation
	result := dao.DB.Where("user_id = ?", userId).Find(&selfEvaluations)
	return selfEvaluations, result.Error
}
