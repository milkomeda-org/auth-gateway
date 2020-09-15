// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import "goa/model"

// Office 组织
type Office struct {
	model.BaseModel
	ParentID uint   `gorm:"not null"` // 上级ID
	Name     string `gorm:"not null"` // 组织名称
	Type     uint   `gorm:"not null"` // 组织类型
}

// OfficeModuleMapping 组织模块关联
type OfficeModuleMapping struct {
	model.BaseModel
	OfficeID uint `gorm:"not null"`
	ModuleID uint `gorm:"not null"`
}

// OfficeRoleMapping 组织角色关联
type OfficeRoleMapping struct {
	model.BaseModel
	OfficeID uint `gorm:"not null"`
	RoleID   uint `gorm:"not null"`
}
