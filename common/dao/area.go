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
