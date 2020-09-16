// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import "goa/model"

// Office 组织
type Office struct {
	model.BaseModel
	ParentID uint   `gorm:"not null;comment:'上级ID'"` // 上级ID
	Name     string `gorm:"not null;comment:'组织名称'"` // 组织名称
	Type     uint   `gorm:"not null;comment:'组织类型'"` // 组织类型
}

// OfficeModuleMapping 组织模块关联
type OfficeModuleMapping struct {
	model.BaseModel
	OfficeID uint `gorm:"not null;comment:'组织ID'"`
	ModuleID uint `gorm:"not null;comment:'模块ID'"`
}

// OfficeRoleMapping 组织角色关联
type OfficeRoleMapping struct {
	model.BaseModel
	OfficeID uint `gorm:"not null;comment:'组织ID'"`
	RoleID   uint `gorm:"not null;comment:'角色ID'"`
}
