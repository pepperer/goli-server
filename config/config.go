package config

import (
	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}
}
