package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
)

// IAreaDAO 定义AreaDAO接口
type IAreaDAO interface {
	Get(id int) (*model.Area, error)
	GetList(pid int) ([]*model.Area, error)
	GetByMergerNameAndLevel(req model.GetByMergerNameAndLevelReq) ([]*model.Area, error)
	GetListByFirstLetter() ([]*model.Area, error)
	GetProvincesWithCities() ([]*model.ProvinceWithCities, error)
	GetAreaIDs(pid int) ([]int, error)
	GetAreaNameById(id int) (string, error)
}

// AreaDAO 实现结构体
type AreaDAO struct {
	DB *gorm.DB // 依赖注入数据库连接
}

func tabName() string {
	return "area"
}

func NewAreaDAO(db *gorm.DB) IAreaDAO {
	return &AreaDAO{DB: db}
}

func (dao *AreaDAO) Get(id int) (*model.Area, error) {
	var area model.Area
	result := dao.DB.Table(tabName()).Where("id = ?", id).First(&area)
	return &area, result.Error
}

func (dao *AreaDAO) GetList(pid int) ([]*model.Area, error) {
	var areas []*model.Area
	result := dao.DB.Table(tabName()).Where("pid = ?", pid).Find(&areas)
	return areas, result.Error
}

// GetAreaNameById 根据 id 获取区域 名称
func (dao *AreaDAO) GetAreaNameById(id int) (string, error) {
	var areaName string
	result := dao.DB.Table(tabName()).Where("id = ?", id).First("shortname", &areaName)
	if result.Error != nil {
		return "", result.Error // 查询出错时返回错误
	}
	return areaName, nil
}

// GetAreaIDs 根据 pid 获取区域 ID 列表
func (dao *AreaDAO) GetAreaIDs(pid int) ([]int, error) {
	// 创建一个空的 int 切片用于存储区域 ID
	var areaIDs []int

	// 执行数据库查询，直接将结果扫描到 areaIDs 切片中
	result := dao.DB.Table(tabName()).Where("pid = ?", pid).Pluck("id", &areaIDs)
	if result.Error != nil {
		return nil, result.Error // 查询出错时返回错误
	}

	// 返回填充好的区域 ID 切片和 nil 错误
	return areaIDs, nil
}

func (dao *AreaDAO) GetByMergerNameAndLevel(req model.GetByMergerNameAndLevelReq) ([]*model.Area, error) {
	var areas []*model.Area
	// 使用LIKE操作符来匹配包含mergerName的记录，并且限制Level为3
	result := dao.DB.Table(tabName()).Where("merger_name LIKE ? AND level = ?", "%"+req.MergerName+"%", req.Level).Find(&areas)
	return areas, result.Error
}

//func (dao *AreaDAO) GetListByFirstLetter(first string) ([]*model.Area, error) {
//	var areas []*model.Area
//	// 使用LIKE查询来匹配首字母，假设first字段存储的就是每个地区名称的首字母
//	result := dao.DB.Where("first = ?", first).Find(&areas)
//	return areas, result.Error
//}

func (dao *AreaDAO) GetListByFirstLetter() ([]*model.Area, error) {
	var areas []*model.Area
	// 按照首字母排序
	result := dao.DB.Table(tabName()).Where("level != ?", 3).Order("first ASC").Find(&areas)
	return areas, result.Error
}

// GetProvincesWithCities 获取所有省份及其下属城市
func (dao *AreaDAO) GetProvincesWithCities() ([]*model.ProvinceWithCities, error) {
	var provinces []*model.Area
	// 首先获取所有省份
	result := dao.DB.Table(tabName()).Where("level = ?", 1).Find(&provinces)
	if result.Error != nil {
		return nil, result.Error
	}

	var provincesWithCities []*model.ProvinceWithCities
	for _, province := range provinces {
		var cities []*model.Area
		// 对每个省份，查询其下属的城市
		result := dao.DB.Table(tabName()).Where("pid = ?", province.Id).Find(&cities)
		if result.Error != nil {
			return nil, result.Error
		}
		// 组装省份及其城市数据
		provinceWithCities := &model.ProvinceWithCities{
			Province: province,
			Cities:   cities,
		}
		provincesWithCities = append(provincesWithCities, provinceWithCities)
	}

	return provincesWithCities, nil
}
