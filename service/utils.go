package service

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

	//用户注册
	ErrorUserExist      = 100
	ErrorEmailExist     = 101
	ErrorFailEncryption = 102
)

var MsgFlags = map[int]string {
	Success:       "ok",
	Unknown:       "Unknown error",
	ParamsError:   "参数错误",
	Error:         "system error",
	ErrorDatabase: "数据库执行错误",

	//用户注册
	ErrorUserExist:      "用户名重复",
	ErrorEmailExist:     "邮箱重复",
	ErrorFailEncryption: "密码加密失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Unknown]
	}
	return msg
}