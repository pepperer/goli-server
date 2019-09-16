package main

import (
	"fmt"
	"goli/config"
	"goli/logger"
	"goli/model"
	"os"
)

func main() {
	// 初始化
	config.Init()
	// 连接数据库
	model.Init(os.Getenv("database"))
	// 连接查询

	var u model.User
	model.DB.First(&u)
	//model.DB.First(&v)
	//fmt.Print(u)
	//
	// Related
	model.DB.Model(&u).Related(&u.Videos).Find(&u.Videos)

	// Association
	//model.DB.Model(&u).Association("Videos").Find(&u.Videos)

	//model.DB.Preload("Videos").First(&u)
	fmt.Println(u)

	var list []model.User
	model.DB.Debug().Preload("Videos").Find(&list)
	logger.Println(list)
}
