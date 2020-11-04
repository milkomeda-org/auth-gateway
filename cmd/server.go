package main

import (
	"github.com/gin-gonic/gin"
	"oa-auth/initializer"
	"oa-auth/initializer/db"
	"oa-auth/router"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// 从本地读取环境变量
	_ = godotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	//组件初始化
	initializer.InitLogger()
	db.InitDB()
	//tools.Migration()
	//系统初始化
	initializer.InitSystem()
	// 装载路由
	r := router.NewRouter()
	_ = r.Run(os.Getenv("SERVER_HOST"))
}
