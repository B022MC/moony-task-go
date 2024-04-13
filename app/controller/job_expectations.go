package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
)

type JobExpectationController struct {
	service service.IJobExpectationService
}

func NewJobExpectationController(service service.IJobExpectationService) *JobExpectationController {
	return &JobExpectationController{
		service: service,
	}
}

func (jec *JobExpectationController) GetJobExpectation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	jobExpectation, err := jec.service.GetJobExpectation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobExpectation)
}

func (jec *JobExpectationController) GetAllJobExpectationsByUserId(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	jobExpectations, err := jec.service.GetAllJobExpectationsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobExpectations)
}

func (jec *JobExpectationController) CreateJobExpectation(c *gin.Context) {
	var jobExpectation model.JobExpectation
	if err := c.ShouldBindJSON(&jobExpectation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := jec.service.CreateJobExpectation(&jobExpectation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (jec *JobExpectationController) UpdateJobExpectation(c *gin.Context) {
	var jobExpectation model.JobExpectation
	if err := c.ShouldBindJSON(&jobExpectation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := jec.service.UpdateJobExpectation(&jobExpectation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (jec *JobExpectationController) DeleteJobExpectation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := jec.service.DeleteJobExpectation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "JobRsp expectation deleted successfully"})
}
