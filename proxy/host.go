// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/12/23.

package proxy

import (
	"auth-gateway/enums"
	"auth-gateway/service/proxy"
	"auth-gateway/util/log"
	"net/http"
	"net/http/httputil"
)

var hosts = make(map[int]*httputil.ReverseProxy)

// InitHost 刷新Host
func InitHost() {
	list := proxy.HostList()
	if nil != list {
		for _, host := range *list {
			if hosts[host.ID] != nil {
				continue
			}
			h :=
				httputil.ReverseProxy{
					Director: func(req *http.Request) {
						// 还原成主机path
						req.URL.Path = req.URL.Path[len(enums.RoutePrefix+enums.ProxyPrefix):]
						req.URL.Scheme = host.URLScheme
						req.URL.Host = host.URLHost
						req.Host = host.Host
					},
					ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {
						log.Error("http: proxy error: %v", err)
						writer.WriteHeader(http.StatusBadGateway)
					},
				}
			hosts[host.ID] = &h
		}
	}
}

// GetHost 获取Host
func GetHost(id int) *httputil.ReverseProxy {
	return hosts[id]
}
