package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
)

type JobsController struct {
	service service.IJobsService
}

func NewJobsController(service service.IJobsService) *JobsController {
	return &JobsController{
		service: service,
	}
}

// GetJob 获取单个兼职信息
func (jc *JobsController) GetJob(c *gin.Context) {
	jobIdStr := c.Param("jobId")
	jobId, err := strconv.ParseInt(jobIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}
	job, err := jc.service.GetJob(jobId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}

// GetAllJobs 获取所有兼职信息
func (jc *JobsController) GetAllJobs(c *gin.Context) {
	// 这里简化了参数的处理，实际应用中可以根据需要获取分页参数等
	comReq := model.ComReq{} // 假设这是通用请求参数，包含分页信息等
	jobs, err := jc.service.GetAllJobs(comReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

// GetRecentJobs 获取最近发布的兼职信息
func (jc *JobsController) GetRecentJobs(c *gin.Context) {
	comReq := model.ComReq{} // 使用通用请求参数
	jobs, err := jc.service.GetRecentJobs(comReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

// GetJobsNearby 获取附近的兼职信息
func (jc *JobsController) GetJobsNearby(c *gin.Context) {
	// 简化了参数处理，实际应用中应从请求中获取经纬度和半径等信息
	lat, lng, radius := 0.0, 0.0, 0 // 假设的经纬度和半径
	comReq := model.ComReq{}        // 使用通用请求参数
	jobs, err := jc.service.GetJobsNearby(lat, lng, radius, comReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

// CreateJob 创建新的兼职信息
func (jc *JobsController) CreateJob(c *gin.Context) {
	var job model.Jobs
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := jc.service.CreateJob(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// UpdateJob 更新兼职信息
func (jc *JobsController) UpdateJob(c *gin.Context) {
	var job model.Jobs
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := jc.service.UpdateJob(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteJob 删除兼职信息
func (jc *JobsController) DeleteJob(c *gin.Context) {
	jobIdStr := c.Param("jobId")
	jobId, err := strconv.ParseInt(jobIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}

	err = jc.service.DeleteJob(jobId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}
