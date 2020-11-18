// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package tools

import (
	"oa-auth/initializer/db"
	"oa-auth/model"
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
