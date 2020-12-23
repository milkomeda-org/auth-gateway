package proxy

import (
	"auth-gateway/cache"
	"auth-gateway/proxy"
	"auth-gateway/serializer"
	proxy2 "auth-gateway/service/proxy"
	"auth-gateway/util/log"

	"github.com/gin-gonic/gin"
)

// RegisterProxy 注册代理
func RegisterProxy(c *gin.Context) {
	var rs proxy2.RegisterService
	if err := c.ShouldBind(&rs); err == nil {
		handle := proxy.GetHost(rs.HostID)
		if nil != handle {
			res := rs.Register()
			cache.AppProxy.Handle(rs.Method, rs.Path, func(context *gin.Context) {
				handle.ServeHTTP(context.Writer, context.Request)
			})
			c.JSON(200, res)
		} else {
			log.Error("host %d is not found", rs.HostID)
			c.JSON(200, serializer.I18Error(err))
		}
	} else {
		c.JSON(200, serializer.I18Error(err))
	}
}

// RegisterProxyHost 注册代理主机
func RegisterProxyHost(c *gin.Context) {
	var rs proxy2.HostRegisterService
	if err := c.ShouldBind(&rs); err == nil {
		res := rs.Register()
		c.JSON(200, res)
		go proxy.InitHost()
	} else {
		c.JSON(200, serializer.I18Error(err))
	}
}

// ViewProxyHost 注册代理主机
func ViewProxyHost(c *gin.Context) {
	c.JSON(200, proxy2.HostList())
}
