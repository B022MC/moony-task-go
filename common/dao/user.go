package dao

import (
	"gorm.io/gorm"
	"moony-task-go/common/model"
	"moony-task-go/core/global"
)

type User struct {
}

func UserInstance() *User {
	return &User{}
}

func (u *User) TableName() string {
	return "user"
}

// Get 获取单个
func (u *User) Get(id int64) (*model.User, error) {
	var user *model.User
	result := global.Db.Table(u.TableName()).Where("id = ?", id).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserList 列表
func (u *User) GetUserList(startTime int64, endTime int64) ([]*model.User, error) {
	var user []*model.User
	tx := global.Db.Table(u.TableName())
	if startTime != 0 {
		tx = tx.Where("create_time>=?", startTime)
	}
	if endTime != 0 {
		tx = tx.Where("create_time<", endTime)
	}
	result := tx.Offset(-1).Limit(-1).Find(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetByPhone 通过手机号获取
func (u *User) GetByPhone(phone string) (*model.User, error) {
	var user *model.User
	result := global.Db.Table(u.TableName()).Where("phone = ?", phone).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetByUnionid 通过微信unionid 获取
func (u *User) GetByUnionid(unionid string) (*model.User, error) {
	var user *model.User
	result := global.Db.Table(u.TableName()).Where("unionid = ?", unionid).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// Create 新增
func (u *User) Create(user *model.User) error {
	result := global.Db.Table(u.TableName()).Create(user)
	return result.Error
}

// Update 更新数据
func (u *User) Update(user *model.User) error {
	result := global.Db.Table(u.TableName()).Save(user)
	return result.Error
}

// Delete 删除
func (u *User) Delete(id int64) error {
	res := global.Db.Table(u.TableName()).Delete(&model.User{}, id)
	return res.Error
}
