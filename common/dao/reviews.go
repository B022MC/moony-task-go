package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model" // 确保这里的路径与你的项目结构相匹配
)

// IReviewDAO 定义了 Review 操作的接口
type IReviewDAO interface {
	Get(reviewId int64) (*model.Review, error)
	GetByJobId(jobId int64) ([]*model.Review, error)
	GetByUserId(userId int64) ([]*model.Review, error)
	Delete(reviewId int64) error
	Create(review *model.Review) (*model.Review, error)
	Update(review *model.Review) (*model.Review, error)
}

type ReviewDAO struct {
	DB *gorm.DB
}

// NewReviewDAO 创建一个新的 ReviewDAO 实例
func NewReviewDAO(db *gorm.DB) IReviewDAO {
	return &ReviewDAO{DB: db}
}

// Get 根据 ID 获取单个 Review
func (dao *ReviewDAO) Get(reviewId int64) (*model.Review, error) {
	var review model.Review
	result := dao.DB.Where("review_id = ?", reviewId).First(&review)
	return &review, result.Error
}

// GetByJobId 根据 JobId 获取 Review 列表
func (dao *ReviewDAO) GetByJobId(jobId int64) ([]*model.Review, error) {
	var reviews []*model.Review
	result := dao.DB.Where("job_id = ?", jobId).Find(&reviews)
	return reviews, result.Error
}

// GetByUserId 根据 UserId 获取 Review 列表
func (dao *ReviewDAO) GetByUserId(userId int64) ([]*model.Review, error) {
	var reviews []*model.Review
	result := dao.DB.Where("user_id = ?", userId).Find(&reviews)
	return reviews, result.Error
}

// Delete 根据 ID 删除 Review
func (dao *ReviewDAO) Delete(reviewId int64) error {
	result := dao.DB.Delete(&model.Review{}, reviewId)
	return result.Error
}

// Create 创建新的 Review
func (dao *ReviewDAO) Create(review *model.Review) (*model.Review, error) {
	result := dao.DB.Create(review)
	return review, result.Error
}

// Update 更新 Review
func (dao *ReviewDAO) Update(review *model.Review) (*model.Review, error) {
	result := dao.DB.Save(review)
	return review, result.Error
}
