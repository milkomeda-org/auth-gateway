package initializer

import (
	"fmt"
	"oa-auth/util"
)

// 初始化日志
func InitLogger() {
	// init logger
	if err := util.InitLogger(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
}
