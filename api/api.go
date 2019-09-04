package api

import (
	"github.com/gin-gonic/gin"
	"goli/bean"
)

func Ping(context *gin.Context) {
	context.JSON(200, bean.BuildResponse("hello world"))
}
