package api

import (
	"github.com/gin-gonic/gin"
	"singo/cache"
	"singo/proxy"
	"singo/service"
)

// RegisterRouter 注册路由
func RegisterRouter(c *gin.Context) {
	var rs service.RouterRegisterService
	if err := c.ShouldBind(&rs); err == nil {
		res := rs.Register()
		c.JSON(200, res)
		go func() {
			cache.AppProxyRouter.Handle(rs.Method, rs.Path, proxy.HostProxy)
		}()
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
