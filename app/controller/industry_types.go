package controller

import (
	"moony-task-go/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/common/model"
)

type IndustryTypeController struct {
	service service.IndustryTypeService
}

func NewIndustryTypeController(service service.IndustryTypeService) *IndustryTypeController {
	return &IndustryTypeController{
		service: service,
	}
}

// AddIndustryType 添加行业类型
func (c *IndustryTypeController) AddIndustryType(ctx *gin.Context) {
	var industry model.IndustryType
	if err := ctx.ShouldBindJSON(&industry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.AddIndustryType(&industry); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Industry type added"})
}

// GetIndustryTypeById 获取指定 ID 的行业类型
func (c *IndustryTypeController) GetIndustryTypeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	industry, err := c.service.GetIndustryTypeById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, industry)
}

// GetAllIndustriesWithCategories 获取所有行业及其关联类别
func (c *IndustryTypeController) GetAllIndustriesWithCategories(ctx *gin.Context) {
	result, err := c.service.GetAllIndustriesWithCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
