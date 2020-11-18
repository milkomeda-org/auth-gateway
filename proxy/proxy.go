package proxy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"oa-auth/util/log"
	"strings"
)

var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		req.URL.Path = strings.Replace(req.URL.Path, "/api/v1/proxy", "", 1)
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:3000"
		req.Host = "localhost"
	},
	ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {
		log.Error("http: proxy error: %v", err)
		writer.WriteHeader(http.StatusBadGateway)
	},
}

// HostProxy 路由代理
func HostProxy(ctx *gin.Context) {
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
