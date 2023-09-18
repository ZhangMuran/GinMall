package dao

import (
	"context"

	"github.com/ZhangMuran/GinMall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(c context.Context) *UserDao {
	return &UserDao{NewDBClient(c)}
}

func NewUserDaobyDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 判断UserName是否已经存在
func (dao *UserDao) ExistOrNotByUserName(userName string) (bool, error) {
	var count int64
	err := dao.DB.Model(&model.User{}).Where("user_name=?", userName).Count(&count).Error
	if err != nil {
		return true, err
	}
	if count != 1 {
		return false, nil
	}
	return true, nil
}

// ExistOrNotByEmail 判断Email是否已经存在
func (dao *UserDao) ExistOrNotByEmail(Email string) (bool, error) {
	var count int64
	err := dao.DB.Model(&model.User{}).Where("email=?", Email).Count(&count).Error
	if err != nil {
		return true, err
	}
	if count != 1 {
		return false, nil
	}
	return true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error;
}