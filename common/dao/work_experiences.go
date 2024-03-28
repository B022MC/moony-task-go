package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type IWorkExperienceDAO interface {
	Get(id int64) (*model.WorkExperience, error)
	GetByUserId(userId int64) ([]*model.WorkExperience, error)
}

type WorkExperienceDAO struct {
	DB *gorm.DB
}

func NewWorkExperienceDAO(db *gorm.DB) IWorkExperienceDAO {
	return &WorkExperienceDAO{DB: db}
}

func (dao *WorkExperienceDAO) Get(id int64) (*model.WorkExperience, error) {
	var workExperience model.WorkExperience
	result := dao.DB.Where("id = ?", id).First(&workExperience)
	return &workExperience, result.Error
}

func (dao *WorkExperienceDAO) GetByUserId(userId int64) ([]*model.WorkExperience, error) {
	var workExperiences []*model.WorkExperience
	result := dao.DB.Where("user_id = ?", userId).Find(&workExperiences)
	return workExperiences, result.Error
}
