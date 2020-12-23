// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package tools

import (
	"auth-gateway/initializer/db"
	"auth-gateway/model"
)

// Migration 自动建表
func Migration() {
	// 自动迁移模式
	db.DB.AutoMigrate(
		&model.Office{},
		&model.Position{},
		&model.Group{},
		&model.Role{},
		&model.ResRelation{},
		&model.RoleRouter{},
		&model.Module{},
		&model.Proxy{},
		&model.User{},
		&model.UserOauth{},
	)
}
