package conf

import (
	_"fmt"
	"strings"

	"github.com/ZhangMuran/GinMall/dao"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string

	DbUser string
	DbPassWord string
	DbHost string
	DbPort string
	DbName string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}

	LoadServer(file)
	LoadMySQL(file)

	// 加载MySQL
	pathRead := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	pathWrite := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	dao.ConnectDb(pathRead, pathWrite)
}

func LoadServer(file *ini.File) {
	AppMode  = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMySQL(file *ini.File) {
	DbUser     = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassword").String()
	DbHost     = file.Section("mysql").Key("DbHost").String()
	DbName     = file.Section("mysql").Key("DbName").String()
	DbPort     = file.Section("mysql").Key("DbPort").String()
}