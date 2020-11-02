// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import "oa-auth/model"

// Position 身份,职位
type Position struct {
	model.BaseModel
	ParentID uint   `gorm:"not null;comment:'上级ID'"` // 上级ID
	OfficeID uint   `gorm:"not null;comment:'组织ID'"` // 组织ID
	Name     string `gorm:"not null;comment:'职位名称'"` // 职位名称
	Code     string `gorm:"not null;comment:'职位编码'"` //职位编码
}

// PositionRoleMapping 职位角色关联
type PositionRoleMapping struct {
	model.BaseModel
	PositionID uint `gorm:"not null;comment:'职位ID'"`
	RoleID     uint `gorm:"not null;comment:'角色ID'"`
}

// PositionModuleMapping 职位模块关联
type PositionModuleMapping struct {
	model.BaseModel
	PositionID uint `gorm:"not null;comment:'职位ID'"`
	ModuleID   uint `gorm:"not null;comment:'模块ID'"`
}
