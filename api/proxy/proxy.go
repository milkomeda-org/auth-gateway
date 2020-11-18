package proxy

import (
	"oa-auth/cache"
	"oa-auth/proxy"
	"oa-auth/serializer"
	proxy2 "oa-auth/service/proxy"

	"github.com/gin-gonic/gin"
)

// RegisterProxy 注册代理
func RegisterProxy(c *gin.Context) {
	var rs proxy2.RegisterService
	if err := c.ShouldBind(&rs); err == nil {
		cache.AppProxy.Handle(rs.Method, rs.Path, proxy.HostProxy)
		res := rs.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.I18Error(err))
	}
}
