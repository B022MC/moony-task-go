package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

// IAreaDAO 定义AreaDAO接口
type IAreaDAO interface {
	Get(id int) (*model.Area, error)
	GetList(pid int) ([]*model.Area, error)
	GetByLevelList(level int) ([]*model.Area, error)
}

// AreaDAO 实现结构体
type AreaDAO struct {
	DB *gorm.DB // 依赖注入数据库连接
}

func NewAreaDAO(db *gorm.DB) IAreaDAO {
	return &AreaDAO{DB: db}
}

func (dao *AreaDAO) Get(id int) (*model.Area, error) {
	var area model.Area
	result := dao.DB.Where("id = ?", id).First(&area)
	return &area, result.Error
}

func (dao *AreaDAO) GetList(pid int) ([]*model.Area, error) {
	var areas []*model.Area
	result := dao.DB.Where("pid = ?", pid).Find(&areas)
	return areas, result.Error
}

func (dao *AreaDAO) GetByLevelList(level int) ([]*model.Area, error) {
	var areas []*model.Area
	result := dao.DB.Where("level = ?", level).Find(&areas)
	return areas, result.Error
}
