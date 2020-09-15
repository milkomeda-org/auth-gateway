package main

import (
	"goa/initializer"
	"goa/server"
	"goa/tools"
	"os"
)

func main() {
	initializer.InitDB()
	initializer.InitLogger()
	//自动迁移建表
	tools.Migration()
	// 装载路由
	r := server.NewRouter()
	_ = r.Run(os.Getenv("SERVER_HOST"))
}
