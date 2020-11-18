package log

import (
	"fmt"
	"oa-auth/util/log"
)

// 初始化日志
func InitLogger() {
	// init log
	if err := log.InitLogger(); err != nil {
		fmt.Printf("init log failed, err:%v\n", err)
		return
	}
}
