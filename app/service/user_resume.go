package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

// IUserResumeService 定义用户简历服务接口
type IUserResumeService interface {
	GetUserResumeByUserId(userId int64) (*model.UserResume, error)
}

// UserResumeService 实现了IUserResumeService接口
type UserResumeService struct{}

// NewUserResumeService 创建UserResumeService的实例
func NewUserResumeService() IUserResumeService {
	return &UserResumeService{}
}

// GetUserResumeByUserId 根据用户ID获取用户简历信息
func (s *UserResumeService) GetUserResumeByUserId(userId int64) (*model.UserResume, error) {
	// 这里使用全局的Db实例来创建UserResumeDAO的实例
	return dao.NewUserResumeDAO(global.Db).GetByUserId(userId)
}
