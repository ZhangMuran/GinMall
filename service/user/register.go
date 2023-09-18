package user

import (
	"context"

	"github.com/ZhangMuran/GinMall/dao"
	"github.com/ZhangMuran/GinMall/model"
	"github.com/ZhangMuran/GinMall/service"
)

//UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname  string `form:"nickname" json:"nickname"`
	UserName  string `form:"user_name" json:"user_name"`
	Password  string `form:"password" json:"password"`
}

func (u UserRegisterService)Register(c context.Context) service.Response {
	// 判断用户名是否重复
	userDao := dao.NewUserDao(c)
	exist, err := userDao.ExistOrNotByUserName(u.UserName)
	if err != nil {
		return service.Response{
			Errno:  service.Error,
			ErrMsg: service.GetMsg(service.Error),
		}
	}
	if exist {
		return service.Response{
			Errno:  service.ErrorUserExist,
			ErrMsg: service.GetMsg(service.ErrorUserExist),
		}
	}

	// 创建用户
	user := model.User{
		Nickname: u.Nickname,
		UserName: u.UserName,
		Status:   model.Active,
	}
	// 给密码加密
	if err := user.SetPassword(u.Password); err != nil {
		return service.Response{
			Errno:  service.ErrorFailEncryption,
			ErrMsg: service.GetMsg(service.ErrorFailEncryption),
		}
	}
	user.Avatar = "https://s3.bmp.ovh/imgs/2023/06/12/a110dd07614a8593.png"
	if err := userDao.CreateUser(&user); err != nil {
		return service.Response{
			Errno:  service.ErrorDatabase,
			ErrMsg: service.GetMsg(service.ErrorDatabase),
		}
	}
	return service.Response{
		Errno: service.Success,
		ErrMsg: service.GetMsg(service.Success),
	}
}