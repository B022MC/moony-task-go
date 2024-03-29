package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

type ISkillCertificateService interface {
	GetSkillCertificate(id int64) (*model.SkillCertificate, error)
	GetAllSkillCertificatesByUserId(userId int64) ([]*model.SkillCertificate, error)
	DeleteSkillCertificate(id int64) error
	CreateSkillCertificate(skillCertificate *model.SkillCertificate) (*model.SkillCertificate, error)
	UpdateSkillCertificate(skillCertificate *model.SkillCertificate) (*model.SkillCertificate, error)
}

type SkillCertificateService struct{}

func NewSkillCertificateService() ISkillCertificateService {
	return &SkillCertificateService{}
}

func (s *SkillCertificateService) GetSkillCertificate(id int64) (*model.SkillCertificate, error) {
	return dao.NewSkillCertificateDAO(global.Db).Get(id)
}

func (s *SkillCertificateService) GetAllSkillCertificatesByUserId(userId int64) ([]*model.SkillCertificate, error) {
	return dao.NewSkillCertificateDAO(global.Db).GetByUserId(userId)
}

func (s *SkillCertificateService) DeleteSkillCertificate(id int64) error {
	return dao.NewSkillCertificateDAO(global.Db).Delete(id)
}

func (s *SkillCertificateService) CreateSkillCertificate(skillCertificate *model.SkillCertificate) (*model.SkillCertificate, error) {
	return dao.NewSkillCertificateDAO(global.Db).Create(skillCertificate)
}

func (s *SkillCertificateService) UpdateSkillCertificate(skillCertificate *model.SkillCertificate) (*model.SkillCertificate, error) {
	return dao.NewSkillCertificateDAO(global.Db).Update(skillCertificate)
}
