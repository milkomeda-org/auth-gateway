package main

import (
	"oa-auth/initializer"
	"oa-auth/server"
	"oa-auth/tools"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// 从本地读取环境变量
	_ = godotenv.Load()
}

func main() {
	initializer.InitDB()
	initializer.InitLogger()
	//自动迁移建表
	tools.Migration()
	// 装载路由
	r := server.NewRouter()
	_ = r.Run(os.Getenv("SERVER_HOST"))
}
