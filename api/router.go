package api

import (
	"oa-auth/cache"
	"oa-auth/proxy"
	"oa-auth/service"

	"github.com/gin-gonic/gin"
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
