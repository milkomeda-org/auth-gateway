package initializer

import (
	"oa-auth/configs"
	"oa-auth/util"
	"os"
)

// InitDB 初始化配置项
func InitDB() {
	// 读取翻译文件
	if err := configs.LoadLocales(os.Getenv("I18N_MAPPINGS_PATH")); err != nil {
		util.Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	Database(os.Getenv("MYSQL_DSN"))
	// 装载Casbin
	CasbinLoader(os.Getenv("MYSQL_DSN"))
	//cache.Redis()
}
