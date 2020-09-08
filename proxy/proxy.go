package proxy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"strings"
)

var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		req.URL.Path = strings.Replace(req.URL.Path, "/api/v1/proxy", "", 1)
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:3000"
		req.Host = "localhost"
	},
}

// HostProxy 路由代理
func HostProxy(ctx *gin.Context) {
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
