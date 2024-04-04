package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
)

type ApplicationController struct {
	service service.IApplicationService
}

func NewApplicationController(service service.IApplicationService) *ApplicationController {
	return &ApplicationController{
		service: service,
	}
}

// GetApplication 获取单个申请信息
func (ac *ApplicationController) GetApplication(c *gin.Context) {
	applicationIdStr := c.Param("applicationId")
	applicationId, err := strconv.ParseInt(applicationIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid application ID"})
		return
	}
	application, err := ac.service.GetApplication(applicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, application)
}

// GetApplicationsByJobId 获取某兼职信息的所有申请
func (ac *ApplicationController) GetApplicationsByJobId(c *gin.Context) {
	jobIdStr := c.Param("jobId")
	jobId, err := strconv.ParseInt(jobIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}
	applications, err := ac.service.GetApplicationsByJobId(jobId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if applications == nil {
		applications = []*model.Application{} // 确保响应不为nil
	}
	c.JSON(http.StatusOK, applications)
}

// GetApplicationsByUserId 获取某用户的所有申请
func (ac *ApplicationController) GetApplicationsByUserId(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	applications, err := ac.service.GetApplicationsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if applications == nil {
		applications = []*model.Application{} // 确保响应不为nil
	}
	c.JSON(http.StatusOK, applications)
}

// CreateApplication 创建新的申请
func (ac *ApplicationController) CreateApplication(c *gin.Context) {
	var application model.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := ac.service.CreateApplication(&application)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// UpdateApplication 更新申请信息
func (ac *ApplicationController) UpdateApplication(c *gin.Context) {
	var application model.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := ac.service.UpdateApplication(&application)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteApplication 删除申请
func (ac *ApplicationController) DeleteApplication(c *gin.Context) {
	applicationIdStr := c.Param("applicationId")
	applicationId, err := strconv.ParseInt(applicationIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid application ID"})
		return
	}

	err = ac.service.DeleteApplication(applicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}
