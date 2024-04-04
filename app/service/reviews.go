package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

// IReviewService 定义了评价服务接口
type IReviewService interface {
	GetReview(reviewId int64) (*model.Review, error)
	GetReviewsByJobId(jobId int64) ([]*model.Review, error)
	GetReviewsByUserId(userId int64) ([]*model.Review, error)
	CreateReview(review *model.Review) (*model.Review, error)
	UpdateReview(review *model.Review) (*model.Review, error)
	DeleteReview(reviewId int64) error
}

type ReviewService struct {
}

// NewReviewService 创建新的评价服务实例
func NewReviewService() IReviewService {
	return &ReviewService{}
}

func (s *ReviewService) GetReview(reviewId int64) (*model.Review, error) {
	return dao.NewReviewDAO(global.Db).Get(reviewId)
}

func (s *ReviewService) GetReviewsByJobId(jobId int64) ([]*model.Review, error) {
	return dao.NewReviewDAO(global.Db).GetByJobId(jobId)
}

func (s *ReviewService) GetReviewsByUserId(userId int64) ([]*model.Review, error) {
	return dao.NewReviewDAO(global.Db).GetByUserId(userId)
}

func (s *ReviewService) CreateReview(review *model.Review) (*model.Review, error) {
	return dao.NewReviewDAO(global.Db).Create(review)
}

func (s *ReviewService) UpdateReview(review *model.Review) (*model.Review, error) {
	return dao.NewReviewDAO(global.Db).Update(review)
}

func (s *ReviewService) DeleteReview(reviewId int64) error {
	return dao.NewReviewDAO(global.Db).Delete(reviewId)
}
