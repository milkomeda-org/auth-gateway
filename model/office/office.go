// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package office

import "oa-auth/model"

// Office 组织
type Office struct {
	model.BaseModel
	ParentID int    `gorm:"not null;comment:'上级ID'"` // 上级ID
	Name     string `gorm:"not null;comment:'组织名称'"` // 组织名称
	Type     int    `gorm:"not null;comment:'组织类型'"` // 组织类型
}

// OfficeRoleMapping 组织角色关联
type OfficeRoleMapping struct {
	model.BaseModel
	OfficeID int `gorm:"not null;comment:'组织ID'"`
	RoleID   int `gorm:"not null;comment:'角色ID'"`
}

// OfficeModuleMapping 组织模块关联
type OfficeModuleMapping struct {
	model.BaseModel
	OfficeID int `gorm:"not null;comment:'组织ID'"`
	ModuleID int `gorm:"not null;comment:'模块ID'"`
}
