package user

import (
	"context"
	"strings"

	"github.com/ZhangMuran/GinMall/conf"
	"github.com/ZhangMuran/GinMall/dao"
	"github.com/ZhangMuran/GinMall/model"
	"github.com/ZhangMuran/GinMall/service/utils"
	"gopkg.in/mail.v2"
)

type UserEmailService struct {
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// 1-绑定邮箱 2-解绑邮箱 3-修改密码
	OperationType uint `json:"operation_type" form:"operation_type"`
}

func (u UserEmailService)SendEmail(c context.Context, uid uint) utils.Response {
	notice := &model.Notice{}
	token, err := utils.GenerateEmailToken(uid, u.OperationType, u.Email, u.Password)
	if err != nil {
		return utils.Response{
			Errno: utils.ErrAuthToken,
			ErrMsg: utils.GetMsg(utils.ErrAuthToken),
		}
	}

	// 组装邮件内容
	noticeDao := dao.NewNoticeDao(c)
	err = noticeDao.GetNoticeById(u.OperationType, notice)
	if err != nil {
		return utils.Response{
			Errno:  utils.ErrorDatabase,
			ErrMsg: utils.GetMsg(utils.ErrorDatabase),
		}
	}
	address := conf.ValidEmail + token
	mailStr := notice.Text
	mailTex := strings.Replace(mailStr, "Email", address, -1)

	// 开始发送邮件
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", u.Email)
	m.SetHeader("Subject", "Gin-Mall UserInfo change")
	m.SetBody("text/html", mailTex)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		return utils.Response{
			Errno: utils.ErrSendEmail,
			ErrMsg: utils.GetMsg(utils.ErrSendEmail),
		}
	}
	return utils.Response{
		Errno: utils.Success,
		ErrMsg: utils.GetMsg(utils.Success),
	}
}