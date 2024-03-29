package controller

import (
	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
	"net/http"
	"strconv"
)

type WorkExperienceController struct {
	service service.IWorkExperienceService
}

func NewWorkExperienceController(service service.IWorkExperienceService) *WorkExperienceController {
	return &WorkExperienceController{
		service: service,
	}
}

// GetWorkExperience @Summary Get Work Experience
// @Description get work experience by ID
// @Tags workExperience
// @Accept  json
// @Produce  json
// @Param id path int true "Work Experience ID"
// @Success 200 {object} model.WorkExperience
// @Router /workExperience/{id} [get]
func (wec *WorkExperienceController) GetWorkExperience(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	workExperience, err := wec.service.GetWorkExperience(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, workExperience)
}

// GetAllWorkExperiencesByUserId handles fetching all work experiences for a specific user
// @Summary Get All Work Experiences by User ID
// @Description get all work experiences for a specific user
// @Tags workExperience
// @Accept  json
// @Produce  json
// @Param userId path int true "User ID"
// @Success 200 {array} model.WorkExperience
// @Router /users/{userId}/workExperiences [get]
func (wec *WorkExperienceController) GetAllWorkExperiencesByUserId(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	workExperiences, err := wec.service.GetAllWorkExperiencesByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if workExperiences == nil {
		workExperiences = []*model.WorkExperience{} // Ensure the response is not nil
	}
	c.JSON(http.StatusOK, workExperiences)
}

// CreateWorkExperience handles the creation of a new work experience
func (wec *WorkExperienceController) CreateWorkExperience(c *gin.Context) {
	var workExperience model.WorkExperience
	if err := c.ShouldBindJSON(&workExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := wec.service.CreateWorkExperience(&workExperience)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// UpdateWorkExperience handles the update of an existing work experience
func (wec *WorkExperienceController) UpdateWorkExperience(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var workExperience model.WorkExperience
	if err := c.ShouldBindJSON(&workExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workExperience.Id = id // Ensure the workExperience ID is set correctly
	result, err := wec.service.UpdateWorkExperience(&workExperience)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteWorkExperience handles the deletion of a work experience
func (wec *WorkExperienceController) DeleteWorkExperience(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = wec.service.DeleteWorkExperience(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Work experience deleted successfully"})
}
