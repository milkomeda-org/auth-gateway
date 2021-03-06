package main

import (
	"auth-gateway/initializer/db"
	"auth-gateway/initializer/log"
	"auth-gateway/initializer/sys"
	proxy2 "auth-gateway/proxy"
	"auth-gateway/router"
	"auth-gateway/tools"
	"github.com/gin-gonic/gin"
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
	log.InitLogger()
	db.InitDB()
	tools.Migration()
	proxy2.InitHost()
	//系统初始化
	sys.InitSystem()
	// 装载路由
	r := router.New()
	_ = r.Run(os.Getenv("SERVER_HOST"))
}
