package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

// IApplicationDAO 定义了 Application 操作的接口
type IApplicationDAO interface {
	Get(applicationId int64) (*model.Application, error)
	GetByJobId(jobId int64) ([]*model.Application, error)
	GetByUserId(userId int64) ([]*model.Application, error)
	Create(application *model.Application) (*model.Application, error)
	Update(application *model.Application) (*model.Application, error)
	Delete(applicationId int64) error
}

type ApplicationDAO struct {
	DB *gorm.DB
}

// NewApplicationDAO 创建一个新的 ApplicationDAO 实例
func NewApplicationDAO(db *gorm.DB) IApplicationDAO {
	return &ApplicationDAO{DB: db}
}

// Get 根据 ID 获取单个 Application
func (dao *ApplicationDAO) Get(applicationId int64) (*model.Application, error) {
	var application model.Application
	result := dao.DB.Where("application_id = ?", applicationId).First(&application)
	return &application, result.Error
}

// GetByJobId 根据 JobId 获取 Application 列表
func (dao *ApplicationDAO) GetByJobId(jobId int64) ([]*model.Application, error) {
	var applications []*model.Application
	result := dao.DB.Where("job_id = ?", jobId).Find(&applications)
	return applications, result.Error
}

// GetByUserId 根据 UserId 获取 Application 列表
func (dao *ApplicationDAO) GetByUserId(userId int64) ([]*model.Application, error) {
	var applications []*model.Application
	result := dao.DB.Where("user_id = ?", userId).Find(&applications)
	return applications, result.Error
}

// Create 创建新的 Application
func (dao *ApplicationDAO) Create(application *model.Application) (*model.Application, error) {
	result := dao.DB.Create(application)
	return application, result.Error
}

// Update 更新 Application
func (dao *ApplicationDAO) Update(application *model.Application) (*model.Application, error) {
	result := dao.DB.Save(application)
	return application, result.Error
}

// Delete 根据 ID 删除 Application
func (dao *ApplicationDAO) Delete(applicationId int64) error {
	result := dao.DB.Delete(&model.Application{}, applicationId)
	return result.Error
}
