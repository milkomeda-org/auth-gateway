// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package tools

import (
	"oa-auth/initializer/db"
	"oa-auth/model/authorization"
	"oa-auth/model/organization"
	"oa-auth/model/resource"
)

// Migration 自动建表
func Migration() {
	// 自动迁移模式
	db.DB.AutoMigrate(
		&organization.Office{},
		&organization.OfficeModuleMapping{},
		&organization.OfficeRoleMapping{},
		&organization.Position{},
		&organization.PositionModuleMapping{},
		&organization.PositionRoleMapping{},
		&organization.Group{},
		&organization.GroupModuleMapping{},
		&organization.GroupRoleMapping{},
		&organization.GroupUserMapping{},
		&authorization.Role{},
		&authorization.RoleModuleMapping{},
		&authorization.RoleRouterMapping{},
		&resource.Module{},
		&resource.Router{},
		&organization.User{},
		&organization.UserModuleMapping{},
		&organization.UserRoleMapping{},
	)
}
