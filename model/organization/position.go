// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import "goa/model"

// Position 身份,职位
type Position struct {
	model.BaseModel
	ParentID uint   `gorm:"not null"` // 上级ID
	OfficeID uint   `gorm:"not null"` // 组织ID
	Name     string `gorm:"not null"` // 职位名称
	Code     string `gorm:"not null"` //职位编码
}

// PositionUserMapping 职位用户关联
type PositionUserMapping struct {
	model.BaseModel
	PositionID uint `gorm:"not null"`
	UserID     uint `gorm:"not null"`
}

// PositionRoleMapping 职位角色关联
type PositionRoleMapping struct {
	model.BaseModel
	PositionID uint `gorm:"not null"`
	RoleID     uint `gorm:"not null"`
}

// PositionModuleMapping 职位模块关联
type PositionModuleMapping struct {
	model.BaseModel
	PositionID uint `gorm:"not null"`
	ModuleID   uint `gorm:"not null"`
}
