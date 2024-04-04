package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
)

type ReviewController struct {
	service service.IReviewService
}

func NewReviewController(service service.IReviewService) *ReviewController {
	return &ReviewController{
		service: service,
	}
}

// GetReview 获取单个评价信息
func (rc *ReviewController) GetReview(c *gin.Context) {
	reviewIdStr := c.Param("reviewId")
	reviewId, err := strconv.ParseInt(reviewIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid review ID"})
		return
	}
	review, err := rc.service.GetReview(reviewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

// GetReviewsByJobId 获取某兼职信息的所有评价
func (rc *ReviewController) GetReviewsByJobId(c *gin.Context) {
	jobIdStr := c.Param("jobId")
	jobId, err := strconv.ParseInt(jobIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}
	reviews, err := rc.service.GetReviewsByJobId(jobId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// GetReviewsByUserId 获取某用户的所有评价
func (rc *ReviewController) GetReviewsByUserId(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	reviews, err := rc.service.GetReviewsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// CreateReview 创建新的评价
func (rc *ReviewController) CreateReview(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := rc.service.CreateReview(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// UpdateReview 更新评价信息
func (rc *ReviewController) UpdateReview(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := rc.service.UpdateReview(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteReview 删除评价
func (rc *ReviewController) DeleteReview(c *gin.Context) {
	reviewIdStr := c.Param("reviewId")
	reviewId, err := strconv.ParseInt(reviewIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid review ID"})
		return
	}

	err = rc.service.DeleteReview(reviewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
