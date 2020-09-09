package main

import (
	"goa/initializer"
	"goa/server"
	"os"
)

func main() {
	initializer.InitDB()
	initializer.InitLogger()
	// 装载路由
	r := server.NewRouter()
	_ = r.Run(os.Getenv("SERVER_HOST"))
}
