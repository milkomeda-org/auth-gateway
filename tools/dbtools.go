// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package tools

import (
	"goa/initializer"
	"goa/model/authorization"
	"goa/model/organization"
	"goa/model/resource"
)

// Migration 自动建表
func Migration() {
	// 自动迁移模式
	initializer.DB.AutoMigrate(
		&organization.Office{},
		&organization.OfficeModuleMapping{},
		&organization.OfficeRoleMapping{},
		&organization.Position{},
		&organization.PositionModuleMapping{},
		&organization.PositionRoleMapping{},
		&organization.PositionUserMapping{},
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
