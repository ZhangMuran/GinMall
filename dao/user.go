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
func (dao *UserDao) ExistOrNotByUserName(userName string) (exist bool, err error) {
	user := &model.User{}
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	if err == gorm.ErrRecordNotFound || user == nil {
		return false, err
	}
	return true, err
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error;
}