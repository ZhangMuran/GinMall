package user

import (
	"context"

	"github.com/ZhangMuran/GinMall/dao"
	"github.com/ZhangMuran/GinMall/model"
	"github.com/ZhangMuran/GinMall/service/utils"
)

type UserUpdateService struct {
	Nickname  string `form:"nick_name" json:"nick_name"`
}

//当前仅包括nickname的修改
func (u UserUpdateService)Update(c context.Context, uid uint) utils.Response {
	user := &model.User{}
	userDao := dao.NewUserDao(c)
	err := userDao.GetUserById(uid, user)
	if err != nil {
		return utils.Response{
			Errno:  utils.ErrorDatabase,
			ErrMsg: utils.GetMsg(utils.ErrorDatabase),
		}
	}

	// 修改昵称nickname
	if u.Nickname != "" {
		user.Nickname = u.Nickname
	}
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
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