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
	GetByMergerNameAndLevel(req model.GetByMergerNameAndLevelReq) ([]*model.Area, error)
	GetListByFirstLetter() (map[string][]*model.Area, error)
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

// GetByMergerNameAndLevel GetAreasByLevel根据名称和等级获取区域列表
func (s *AreaService) GetByMergerNameAndLevel(req model.GetByMergerNameAndLevelReq) ([]*model.Area, error) {
	areaDao := dao.NewAreaDAO(global.Db)
	return areaDao.GetByMergerNameAndLevel(req)
}

func (s *AreaService) GetListByFirstLetter() (map[string][]*model.Area, error) {
	areaDao := dao.NewAreaDAO(global.Db)
	areas, err := areaDao.GetListByFirstLetter() // 假设这个方法现在返回所有地区
	if err != nil {
		return nil, err
	}

	// 创建一个映射来按首字母组织地区
	areaMap := make(map[string][]*model.Area)
	for _, area := range areas {
		firstLetter := area.First // 假设这是首字母字段
		if _, exists := areaMap[firstLetter]; !exists {
			areaMap[firstLetter] = []*model.Area{} // 如果键不存在，则初始化切片
		}
		areaMap[firstLetter] = append(areaMap[firstLetter], area) // 将地区添加到对应的首字母下
	}

	return areaMap, nil
}
