package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// RegisterRouter 注册路由
func RegisterRouter(c *gin.Context) {
	var service service.RouterRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
