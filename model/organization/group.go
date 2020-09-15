// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import "goa/model"

// Group 组
type Group struct {
	model.BaseModel
	Name string `gorm:"not null"`
	Code string `gorm:"not null"`
}

// GroupUserMapping 用户组关联
type GroupUserMapping struct {
	model.BaseModel
	GroupID uint `gorm:"not null"`
	UserID  uint `gorm:"not null"`
}

// GroupRoleMapping 组权限关联
type GroupRoleMapping struct {
	model.BaseModel
	GroupID uint `gorm:"not null"`
	RoleID  uint `gorm:"not null"`
}

// GroupModuleMapping 组模块关联
type GroupModuleMapping struct {
	model.BaseModel
	GroupID  uint `gorm:"not null"`
	ModuleID uint `gorm:"not null"`
}
