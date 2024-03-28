package dao

import (
	"errors"
	"gorm.io/gorm"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

type Experiment struct {
}

func ConfigDaoInstance() *Experiment {
	return &Experiment{}
}

func (e *Experiment) TableName() string {
	return "bus_experiment"
}

// GetAll 查询所有配置
func (e *Experiment) GetAll(appId int64) ([]*model.Experiment, error) {
	var conf []*model.Experiment
	tx := global.Db.Table(e.TableName())
	if appId != 0 {
		tx = tx.Where("app_id=?", appId)
	}
	tx.Where("status = ?", model.ConfigStatusNormal)

	result := tx.Offset(0).Limit(100).Order("id desc").Find(&conf)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return conf, nil
}
