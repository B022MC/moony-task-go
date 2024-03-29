package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type ISelfEvaluationDAO interface {
	Get(id int64) (*model.SelfEvaluation, error)
	GetByUserId(userId int64) ([]*model.SelfEvaluation, error)
	Delete(id int64) error
	Create(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error)
	Update(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error)
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
func (dao *SelfEvaluationDAO) Delete(id int64) error {
	result := dao.DB.Delete(&model.SelfEvaluation{}, id)
	return result.Error
}

func (dao *SelfEvaluationDAO) Create(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error) {
	result := dao.DB.Create(selfEvaluation)
	return selfEvaluation, result.Error
}

func (dao *SelfEvaluationDAO) Update(selfEvaluation *model.SelfEvaluation) (*model.SelfEvaluation, error) {
	result := dao.DB.Save(selfEvaluation)
	return selfEvaluation, result.Error
}
