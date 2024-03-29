package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

// IUserResumeDAO 定义对UserResume的数据访问接口
type IUserResumeDAO interface {
	GetByUserId(userId int64) (*model.UserResume, error)
}

// UserResumeDAO 实现IUserResumeDAO接口
type UserResumeDAO struct {
	DB *gorm.DB
}

// NewUserResumeDAO 创建UserResumeDAO的实例
func NewUserResumeDAO(db *gorm.DB) IUserResumeDAO {
	return &UserResumeDAO{DB: db}
}

// GetByUserId 根据用户ID查询用户简历信息
func (dao *UserResumeDAO) GetByUserId(userId int64) (*model.UserResume, error) {
	var userResume model.UserResume
	// 这里使用GORM的原生SQL查询方法来查询视图
	result := dao.DB.Raw("SELECT * FROM vw_user_resume WHERE user_id = ?", userId).Scan(&userResume)
	return &userResume, result.Error
}
