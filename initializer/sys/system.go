// Copyright The ef Co. All rights reserved.
// Created by vinson on 2020/11/4.

package sys

import (
	"auth-gateway/service/role"
	"auth-gateway/service/user"
	"fmt"
)

func InitSystem() {
	initRootUser()
	initRootRole()
}

func initRootUser() {
	// user
	if !user.Exists("admin") {
		var us user.RegisterService
		us.UserName = "admin"
		us.Nickname = "admin"
		us.Password = "P@ssW0rd"
		us.PasswordConfirm = "P@ssW0rd"
		res := us.Register()
		fmt.Println(res)
	}
	// mapping
}

func initRootRole() {
	// role
	if !role.Exists("root") {
		var rs role.CreateService
		rs.Alias = "root"
		res := rs.Execute()
		fmt.Println(res)
	}
}
