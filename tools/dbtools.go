// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package tools

import (
	"oa-auth/initializer/db"
	"oa-auth/model/group"
	"oa-auth/model/office"
	"oa-auth/model/position"
	"oa-auth/model/resource"
	"oa-auth/model/role"
	"oa-auth/model/user"
)

// Migration 自动建表
func Migration() {
	// 自动迁移模式
	db.DB.AutoMigrate(
		&office.Office{},
		&office.OfficeModuleMapping{},
		&office.OfficeRoleMapping{},
		&position.Position{},
		&position.PositionModuleMapping{},
		&position.PositionRoleMapping{},
		&group.Group{},
		&group.GroupModuleMapping{},
		&group.GroupRoleMapping{},
		&group.GroupUserMapping{},
		&role.Role{},
		&role.RoleModuleMapping{},
		&role.RoleRouterMapping{},
		&resource.Module{},
		&resource.Router{},
		&user.User{},
		&user.UserModuleMapping{},
		&user.UserRoleMapping{},
	)
}
