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

	//  中间件
	//engine.Use(logger())

	engine.GET("/", api.Ping)
	userRouter := engine.Group("/user")
	{
		userRouter.GET("/list", api.Ping)
	}
	_ = engine.Run(":80")
}

func logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("custom", "custom")
		context.Next()
	}
}
