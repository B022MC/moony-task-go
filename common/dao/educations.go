package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type IEducationDAO interface {
	Get(id int64) (*model.Education, error)
	GetByUserId(userId int64) ([]*model.Education, error)
}

type EducationDAO struct {
	DB *gorm.DB
}

func NewEducationDAO(db *gorm.DB) IEducationDAO {
	return &EducationDAO{DB: db}
}

func (dao *EducationDAO) Get(id int64) (*model.Education, error) {
	var education model.Education
	result := dao.DB.Where("id = ?", id).First(&education)
	return &education, result.Error
}

func (dao *EducationDAO) GetByUserId(userId int64) ([]*model.Education, error) {
	var educations []*model.Education
	result := dao.DB.Where("user_id = ?", userId).Find(&educations)
	return educations, result.Error
}
