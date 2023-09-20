package user

import (
	"context"

	"github.com/ZhangMuran/GinMall/dao"
	"github.com/ZhangMuran/GinMall/model"
	"github.com/ZhangMuran/GinMall/service/utils"
)

//UserRegisterService 管理用户登录服务
type UserLoginService struct {
	UserName  string `form:"user_name" json:"user_name"`
	Password  string `form:"password" json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
	UserInfo
}

type UserInfo struct {
	ID       uint   `json:"id"`
    UserName string `json:"user_name"`
    NickName string `json:"nick_name"`
    Email    string `json:"email"`
    Status   string `json:"status"`
    Avatar   string `json:"avatar"`
    CreateAt int64  `json:"create_at"`
}

func (u *UserLoginService)Login(ctx context.Context) utils.Response {
	user := &model.User{}
	userDao := dao.NewUserDao(ctx)

	// 检查用户是否存在
	exist, err := userDao.ExistOrNotByUserName(u.UserName, user)
	if err != nil {
		return utils.Response{
			Errno:  utils.ErrorDatabase,
			ErrMsg: utils.GetMsg(utils.ErrorDatabase),
		}
	}
	if !exist {
		return utils.Response{
			Errno:  utils.ErrorUserNotFound,
			ErrMsg: utils.GetMsg(utils.ErrorUserNotFound),
			Data: "用户不存在，请先注册",
		}
	}

	// 判断密码是否正确
	if !user.CheckPassword(u.Password) {
		return utils.Response{
			Errno:  utils.ErrorPassword,
			ErrMsg: utils.GetMsg(utils.ErrorPassword),
			Data: "密码错误，请检查后重新尝试",
		}
	}

	//成功登陆了，签发token
	token, err := utils.GenerateToken(user.ID, u.UserName, 0)
	if err != nil {
		return utils.Response{
			Errno: utils.ErrAuthToken,
			ErrMsg: utils.GetMsg(utils.ErrAuthToken),
		}
	}

	return utils.Response{
		Errno: utils.Success,
		ErrMsg: utils.GetMsg(utils.Success),
		Data: UserLoginResponse{
			Token: token,
			UserInfo: UserInfo{
				ID: user.ID,
				UserName: user.UserName,
				NickName: user.Nickname,
				Email: user.Email,
				Status: user.Status,
				Avatar: user.Avatar,
				CreateAt: user.CreatedAt.Unix(),
			},
		},
	}
}