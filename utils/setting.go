package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode   string
	HttpPoort string
	JwtKey    string
	Db        string
	DbHost    string
	DbPort    string
	DbUser    string
	DbPasswd  string
	DbName    string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPoort = file.Section("server").Key("HttpPoort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("lisne92LI4")

}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("192.168.137.43")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPasswd = file.Section("database").Key("DbPasswd").MustString("qwe123qaz")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

}
