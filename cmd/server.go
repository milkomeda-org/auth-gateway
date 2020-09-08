package main

import (
	"goa/initializer"
	"goa/server"
	"strconv"
)

func main() {
	// 装载路由
	r := server.NewRouter()
	_ = r.Run(initializer.Root.Server.Host + ":" + strconv.Itoa(int(initializer.Root.Server.Port)))
}
