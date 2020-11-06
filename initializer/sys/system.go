// Copyright The ef Co. All rights reserved.
// Created by vinson on 2020/11/4.

package sys

import (
	"fmt"
	"oa-auth/service/user"
)

func InitSystem() {
	initRootUser()
}

func initRootUser() {
	if !user.Exists("admin") {
		var us user.RegisterService
		us.UserName = "admin"
		us.Nickname = "admin"
		us.Password = "P@ssW0rd"
		us.PasswordConfirm = "P@ssW0rd"
		res := us.Register()
		fmt.Println(res)
	}
}
