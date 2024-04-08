package controller

import (
	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
	"net/http"
)

type LocationController struct {
	locationService service.LocationService
}

func NewLocationController() *LocationController {
	return &LocationController{}
}

func (lc *LocationController) ReverseGeocode(c *gin.Context) {
	var location model.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := lc.locationService.GetAddressByCoordinates(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 如果响应为空，返回空响应
	if response == nil {
		c.JSON(http.StatusOK, gin.H{"result": nil})
		return
	}

	// 否则返回整个响应对象
	c.JSON(http.StatusOK, gin.H{"result": response.Result})
}
