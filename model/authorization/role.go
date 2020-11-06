package authorization

import "oa-auth/model"

// Role 角色模型
type Role struct {
	model.BaseModel
	Alias string `gorm:"not null comment:'名称'"`
}

// RoleRouterMapping 角色路由关联
type RoleRouterMapping struct {
	model.BaseModel
	RoleID   int `gorm:"not null comment:'角色ID'"`
	RouterID int `gorm:"not null comment:'路由ID'"`
}

// RoleModuleMapping 角色模块关联
type RoleModuleMapping struct {
	model.BaseModel
	RoleID   int `gorm:"not null comment:'角色ID'"`
	ModuleID int `gorm:"not null comment:'模块ID'"`
}
