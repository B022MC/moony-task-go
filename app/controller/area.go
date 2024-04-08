package controller

import (
	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"net/http"
	"strconv"
)

// AreaController 定义区域控制器结构体
type AreaController struct {
	areaService service.IAreaService
}

// NewAreaController 创建区域控制器实例
func NewAreaController(areaService service.IAreaService) *AreaController {
	return &AreaController{
		areaService: areaService,
	}
}

// GetArea 获取单个区域的信息
func (ac *AreaController) GetArea(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	area, err := ac.areaService.GetArea(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, area)
}

// GetSubAreas 根据父ID获取子区域列表
func (ac *AreaController) GetSubAreas(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	areas, err := ac.areaService.GetSubAreas(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, areas)
}

// GetAreasByLevel 根据等级获取区域列表
func (ac *AreaController) GetAreasByLevel(c *gin.Context) {
	level, _ := strconv.Atoi(c.Query("level"))
	areas, err := ac.areaService.GetAreasByLevel(level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, areas)
}

// GetListByFirstLetter 根据首字母获取区域列表
func (ac *AreaController) GetListByFirstLetter(c *gin.Context) {
	//first := cast.ToString(c.Query("first"))
	areas, err := ac.areaService.GetListByFirstLetter()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, areas)
}
