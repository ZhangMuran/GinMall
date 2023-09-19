package user

import (
	"context"

	"github.com/ZhangMuran/GinMall/dao"
	"github.com/ZhangMuran/GinMall/model"
	"github.com/ZhangMuran/GinMall/service/utils"
)

//UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname  string `form:"nick_name" json:"nick_name"`
	UserName  string `form:"user_name" json:"user_name"`
	Password  string `form:"password" json:"password"`
	Email     string `form:"email" json:"email"`
}

func (u UserRegisterService)Register(c context.Context) utils.Response {
	user := &model.User{}
	// 判断用户名是否重复
	userDao := dao.NewUserDao(c)
	exist, err := userDao.ExistOrNotByUserName(u.UserName, user)
	if err != nil {
		return utils.Response{
			Errno:  utils.ErrorDatabase,
			ErrMsg: utils.GetMsg(utils.ErrorDatabase),
		}
	}
	if exist {
		return utils.Response{
			Errno:  utils.ErrorUserExist,
			ErrMsg: utils.GetMsg(utils.ErrorUserExist),
		}
	}

	// 判断邮箱是否重复
	exist, err = userDao.ExistOrNotByEmail(u.Email)
	if err != nil {
		return utils.Response{
			Errno:  utils.ErrorDatabase,
			ErrMsg: utils.GetMsg(utils.ErrorDatabase),
		}
	}
	if exist {
		return utils.Response{
			Errno:  utils.ErrorEmailExist,
			ErrMsg: utils.GetMsg(utils.ErrorEmailExist),
		}
	}

	// 创建用户
	user = &model.User{
		Nickname: u.Nickname,
		UserName: u.UserName,
		Email:    u.Email,
		Status:   model.Active,
	}
	// 给密码加密
	if err := user.SetPassword(u.Password); err != nil {
		return utils.Response{
			Errno:  utils.ErrorFailEncryption,
			ErrMsg: utils.GetMsg(utils.ErrorFailEncryption),
		}
	}
	user.Avatar = "https://s3.bmp.ovh/imgs/2023/06/12/a110dd07614a8593.png"
	if err := userDao.CreateUser(user); err != nil {
		return utils.Response{
			Errno:  utils.ErrorDatabase,
			ErrMsg: utils.GetMsg(utils.ErrorDatabase),
		}
	}
	return utils.Response{
		Errno: utils.Success,
		ErrMsg: utils.GetMsg(utils.Success),
	}
}