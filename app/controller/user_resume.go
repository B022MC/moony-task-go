package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
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
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}
	userResume, err := urc.service.GetUserResumeByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userResume)
}
