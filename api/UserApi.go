package api

import (
	"github.com/gin-gonic/gin"
	"goli/response"
	"goli/service"
)

// UserRegister 用户注册接口
func Register(c *gin.Context) {
	var params RegisterParams
	if err := c.ShouldBind(&params); err == nil {
		if user, err := service.Register(params); err != nil {
			c.JSON(200, response.BuildErrorResponse(err))
		} else {
			res := response.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, response.BuildErrorResponse(err))
	}
}
