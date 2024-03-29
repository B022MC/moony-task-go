package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
)

type SkillCertificateController struct {
	service service.ISkillCertificateService
}

// NewSkillCertificateController 创建新的技能证书控制器实例
func NewSkillCertificateController(service service.ISkillCertificateService) *SkillCertificateController {
	return &SkillCertificateController{
		service: service,
	}
}

// GetSkillCertificate 获取单个技能证书信息
func (sc *SkillCertificateController) GetSkillCertificate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	skillCertificate, err := sc.service.GetSkillCertificate(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, skillCertificate)
}

// GetAllSkillCertificatesByUserId 获取用户的所有技能证书信息
func (sc *SkillCertificateController) GetAllSkillCertificatesByUserId(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}
	skillCertificates, err := sc.service.GetAllSkillCertificatesByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if skillCertificates == nil {
		skillCertificates = []*model.SkillCertificate{} // 确保响应不为nil
	}
	c.JSON(http.StatusOK, skillCertificates)
}

// CreateSkillCertificate 创建技能证书信息
func (sc *SkillCertificateController) CreateSkillCertificate(c *gin.Context) {
	var skillCertificate model.SkillCertificate
	if err := c.ShouldBindJSON(&skillCertificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := sc.service.CreateSkillCertificate(&skillCertificate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

// UpdateSkillCertificate 更新技能证书信息
func (sc *SkillCertificateController) UpdateSkillCertificate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	var skillCertificate model.SkillCertificate
	if err := c.ShouldBindJSON(&skillCertificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skillCertificate.Id = id // 确保ID正确设置
	result, err := sc.service.UpdateSkillCertificate(&skillCertificate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteSkillCertificate 删除技能证书信息
func (sc *SkillCertificateController) DeleteSkillCertificate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	err = sc.service.DeleteSkillCertificate(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "技能证书删除成功"})
}
