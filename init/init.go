package init

import (
	"goa/cache"
	"goa/configs"
	"goa/model"
	"goa/util"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := configs.LoadLocales("configs/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// 装载Casbin
	model.CasbinLoader(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
