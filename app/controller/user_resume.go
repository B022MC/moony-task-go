package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"moony-task-go/app/service"
	"net/http"
)

type UserResumeController struct {
	service service.IUserResumeService
}

func NewUserResumeController(service service.IUserResumeService) *UserResumeController {
	return &UserResumeController{
		service: service,
	}
}

// GetUserResumeByUserId 获取用户的简历信息
// @Summary Get User Resume by User ID
// @Description get user resume by user ID
// @Tags userResume
// @Accept  json
// @Produce  json
// @Param userId path int true "User ID"
// @Success 200 {object} model.UserResume
// @Router /users/{userId}/resume [get]
func (urc *UserResumeController) GetUserResumeByUserId(c *gin.Context) {
	userIdStr := c.Query("userId")
	userId := cast.ToInt64(userIdStr)
	userResume, err := urc.service.GetUserResumeByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userResume)
}
