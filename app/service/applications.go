package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

// IApplicationService 定义了申请服务接口
type IApplicationService interface {
	GetApplication(applicationId int64) (*model.Application, error)               // 根据申请ID获取单个申请
	GetApplicationsByJobId(jobId int64) ([]*model.Application, error)             // 根据兼职信息ID获取申请列表
	GetApplicationsByUserId(userId int64) ([]*model.Application, error)           // 根据用户ID获取申请列表
	CreateApplication(application *model.Application) (*model.Application, error) // 创建新申请
	UpdateApplication(application *model.Application) (*model.Application, error) // 更新申请信息
	DeleteApplication(applicationId int64) error                                  // 删除申请
}

type ApplicationService struct {
}

func NewApplicationService() IApplicationService {
	return &ApplicationService{}
}

// GetApplication 通过ID检索单个申请
func (s *ApplicationService) GetApplication(applicationId int64) (*model.Application, error) {
	return dao.NewApplicationDAO(global.Db).Get(applicationId)
}

// GetApplicationsByJobId 通过兼职信息ID获取申请列表
func (s *ApplicationService) GetApplicationsByJobId(jobId int64) ([]*model.Application, error) {
	return dao.NewApplicationDAO(global.Db).GetByJobId(jobId)
}

// GetApplicationsByUserId 通过用户ID获取申请列表
func (s *ApplicationService) GetApplicationsByUserId(userId int64) ([]*model.Application, error) {
	return dao.NewApplicationDAO(global.Db).GetByUserId(userId)
}

// CreateApplication 创建一个新的申请
func (s *ApplicationService) CreateApplication(application *model.Application) (*model.Application, error) {
	return dao.NewApplicationDAO(global.Db).Create(application)
}

// UpdateApplication 更新一个已存在的申请
func (s *ApplicationService) UpdateApplication(application *model.Application) (*model.Application, error) {
	return dao.NewApplicationDAO(global.Db).Update(application)
}

// DeleteApplication 通过ID删除一个申请
func (s *ApplicationService) DeleteApplication(applicationId int64) error {
	return dao.NewApplicationDAO(global.Db).Delete(applicationId)
}
