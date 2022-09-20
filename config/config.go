package config

import (
	"os"
	"unicode/utf8"

	"gorm.io/gorm"
)

type configData struct {
	DB     mysqlConfig  `json:"mysql"`
	Server serverConfig `json:"server"`
}

var (
	ExecPath     string
	JsonData     configData
	ServerConfig serverConfig
	DB           *gorm.DB
)

func initPath() {
	// 目录路径分隔符，这里为\\
	sep := string(os.PathSeparator)
	// 执行命令时所在目录
	ExecPath, _ = os.Getwd()
	length := utf8.RuneCountInString(ExecPath)
	lastchar := ExecPath[length-1:]
	if lastchar != sep {
		ExecPath = ExecPath + sep
	}
}
