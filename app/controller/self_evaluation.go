package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
)

// SelfEvaluationController struct to handle self evaluation requests
type SelfEvaluationController struct {
	service service.ISelfEvaluationService
}

// NewSelfEvaluationController creates a new instance of SelfEvaluationController
func NewSelfEvaluationController(service service.ISelfEvaluationService) *SelfEvaluationController {
	return &SelfEvaluationController{
		service: service,
	}
}

// GetSelfEvaluation handles GET requests to fetch a single self evaluation by its ID
func (sec *SelfEvaluationController) GetSelfEvaluation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	selfEvaluation, err := sec.service.GetSelfEvaluation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, selfEvaluation)
}

// GetAllSelfEvaluationsByUserId handles GET requests to list all self evaluations by a specific user
func (sec *SelfEvaluationController) GetAllSelfEvaluationsByUserId(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	selfEvaluations, err := sec.service.GetAllSelfEvaluationsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if selfEvaluations == nil {
		selfEvaluations = []*model.SelfEvaluation{} // Ensure the response is not nil
	}

	c.JSON(http.StatusOK, selfEvaluations)
}

// CreateSelfEvaluation handles POST requests to create a new self evaluation
func (sec *SelfEvaluationController) CreateSelfEvaluation(c *gin.Context) {
	var selfEvaluation model.SelfEvaluation
	if err := c.ShouldBindJSON(&selfEvaluation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := sec.service.CreateSelfEvaluation(&selfEvaluation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// UpdateSelfEvaluation handles PUT requests to update an existing self evaluation
func (sec *SelfEvaluationController) UpdateSelfEvaluation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var selfEvaluation model.SelfEvaluation
	if err := c.ShouldBindJSON(&selfEvaluation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	selfEvaluation.Id = id // Ensure the ID is correctly set
	result, err := sec.service.UpdateSelfEvaluation(&selfEvaluation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteSelfEvaluation handles DELETE requests to remove a self evaluation
func (sec *SelfEvaluationController) DeleteSelfEvaluation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = sec.service.DeleteSelfEvaluation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Self evaluation deleted successfully"})
}
