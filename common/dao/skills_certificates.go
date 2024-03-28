package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

type ISkillCertificateDAO interface {
	Get(id int64) (*model.SkillCertificate, error)
	GetByUserId(userId int64) ([]*model.SkillCertificate, error)
}

type SkillCertificateDAO struct {
	DB *gorm.DB
}

func NewSkillCertificateDAO(db *gorm.DB) ISkillCertificateDAO {
	return &SkillCertificateDAO{DB: db}
}

func (dao *SkillCertificateDAO) Get(id int64) (*model.SkillCertificate, error) {
	var skillCertificate model.SkillCertificate
	result := dao.DB.Where("id = ?", id).First(&skillCertificate)
	return &skillCertificate, result.Error
}

func (dao *SkillCertificateDAO) GetByUserId(userId int64) ([]*model.SkillCertificate, error) {
	var skillCertificates []*model.SkillCertificate
	result := dao.DB.Where("user_id = ?", userId).Find(&skillCertificates)
	return skillCertificates, result.Error
}
