package authorization

import "goa/model"

// Role 角色模型
type Role struct {
	model.BaseModel
	Alias string `gorm:"not null"`
}

// RoleRouterMapping 角色路由关联
type RoleRouterMapping struct {
	model.BaseModel
	RoleID   uint `gorm:"not null"`
	RouterID uint `gorm:"not null"`
}

// RoleModuleMapping 角色模块关联
type RoleModuleMapping struct {
	model.BaseModel
	RoleID   uint `gorm:"not null"`
	ModuleID uint `gorm:"not null"`
}
