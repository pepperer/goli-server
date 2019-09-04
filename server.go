package main

import (
	"github.com/gin-gonic/gin"
	"goli/api"
	"goli/config"
	"goli/model"
	"os"
)

func main() {

	config.Init()

	model.Init(os.Getenv("DataBase"))

	engine := gin.Default()
	engine.GET("/", api.Ping)
	userRouter := engine.Group("/user")
	{
		userRouter.GET("/list", api.Ping)
	}
	_ = engine.Run(":80")
}
