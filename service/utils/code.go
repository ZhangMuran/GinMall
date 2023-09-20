package utils

type Response struct {
	Errno  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`   
}

const (
	Success       = 0
	Unknown       = 1
	ParamsError   = 2
	Error         = 3
	ErrorDatabase = 4

	//用户
	ErrorUserExist      = 100
	ErrorEmailExist     = 101
	ErrorFailEncryption = 102
	ErrorUserNotFound   = 103
	ErrorPassword       = 104
	ErrAuthToken        = 105
	ErrTokenTimeOut     = 106
	ErrSendEmail        = 107
)

var MsgFlags = map[int]string {
	Success:       "ok",
	Unknown:       "Unknown error",
	ParamsError:   "参数错误",
	Error:         "system error",
	ErrorDatabase: "数据库执行错误",

	//用户
	ErrorUserExist:      "用户名重复",
	ErrorEmailExist:     "邮箱重复",
	ErrorFailEncryption: "密码加密失败",
	ErrorUserNotFound:   "未找到对应用户",
	ErrorPassword:       "登录密码错误",
	ErrAuthToken:        "签发Token失败",
	ErrTokenTimeOut:     "token 过期",
	ErrSendEmail:        "邮件发送失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Unknown]
	}
	return msg
}