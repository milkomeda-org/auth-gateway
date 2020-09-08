package main

import (
	"goa/init"
	"goa/server"
	"os"
)

func main() {
	// 从配置文件读取配置
	init.Init()

	// 装载路由
	r := server.NewRouter()
	_ = r.Run(os.Getenv("SERVER_HOST"))
}
