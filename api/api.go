package api

import (
	"github.com/gin-gonic/gin"
	"goli/response"
)

func Ping(context *gin.Context) {
	context.JSON(200, response.BuildResponse("hello world"))
}

