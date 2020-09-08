package initializer

import (
	"github.com/joho/godotenv"
	"goa/cache"
	"goa/configs"
	"goa/model"
	"goa/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Root *configs.Root

// Init 初始化配置项
func init() {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/configs/application.yaml")

	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	root := configs.Root{}
	err = yaml.Unmarshal(bytes, &root)
	Root = &root
	if err != nil {
		util.Log().Error("error: %v", err)
	}
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := configs.LoadLocales(os.Getenv("I18N_MAPPINGS_PATH")); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// 装载Casbin
	model.CasbinLoader(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
