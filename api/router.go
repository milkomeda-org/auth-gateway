package api

import (
	"github.com/gin-gonic/gin"
	"singo/cache"
	"singo/proxy"
	"singo/service"
)

// RegisterRouter 注册路由
func RegisterRouter(c *gin.Context) {
	var service service.RouterRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
		go func() {
			cache.AppProxyRouter.Handle(service.Method, service.Path, proxy.HostProxy)
		}()
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
