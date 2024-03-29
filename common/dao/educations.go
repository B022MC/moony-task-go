package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type IEducationDAO interface {
	Get(id int64) (*model.Education, error)
	GetByUserId(userId int64) ([]*model.Education, error)
	Delete(id int64) error
	Create(education *model.Education) (*model.Education, error)
	Update(education *model.Education) (*model.Education, error)
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
func (dao *EducationDAO) Delete(id int64) error {
	result := dao.DB.Delete(&model.Education{}, id)
	return result.Error
}

func (dao *EducationDAO) Create(education *model.Education) (*model.Education, error) {
	result := dao.DB.Create(education)
	return education, result.Error
}

func (dao *EducationDAO) Update(education *model.Education) (*model.Education, error) {
	result := dao.DB.Save(education)
	return education, result.Error
}
