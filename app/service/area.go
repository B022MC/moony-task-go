package service

import (
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

// IAreaService 定义区域服务接口
type IAreaService interface {
	GetArea(id int) (*model.Area, error)
	GetSubAreas(pid int) ([]*model.Area, error)
	GetAreasByLevel(level int) ([]*model.Area, error)
}

// AreaService AreaService结构体
type AreaService struct {
}

// NewAreaService NewAreaService构造函数
func NewAreaService() IAreaService {
	return &AreaService{}
}

// GetArea GetArea根据ID获取单个区域
func (s *AreaService) GetArea(id int) (*model.Area, error) {
	areaDao := dao.NewAreaDAO(global.Db) // 使用全局DB实例
	return areaDao.Get(id)
}

// GetSubAreas GetSubAreas根据父ID获取子区域列表
func (s *AreaService) GetSubAreas(pid int) ([]*model.Area, error) {
	areaDao := dao.NewAreaDAO(global.Db)
	return areaDao.GetList(pid)
}

// GetAreasByLevel GetAreasByLevel根据等级获取区域列表
func (s *AreaService) GetAreasByLevel(level int) ([]*model.Area, error) {
	areaDao := dao.NewAreaDAO(global.Db)
	return areaDao.GetByLevelList(level)
}
