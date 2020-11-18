package main

import (
	"github.com/gin-gonic/gin"
	"oa-auth/initializer/db"
	"oa-auth/initializer/log"
	"oa-auth/initializer/sys"
	"oa-auth/router"
	"oa-auth/tools"
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
	//系统初始化
	sys.InitSystem()
	// 装载路由
	r := router.NewRouter()
	_ = r.Run(os.Getenv("SERVER_HOST"))
}
