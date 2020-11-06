// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package rrt

// ResRelationType 资源关联类型
type ResRelationType int

const (
	// RoleModule 角色-模块
	RoleModule ResRelationType = iota

	//GroupModule 组-模块
	GroupModule
	//GroupRole 组-角色
	GroupRole

	//OfficeModule 组织-模块
	OfficeModule
	//OfficeRole 组织-角色
	OfficeRole

	//PositionModule 身份-模块
	PositionModule
	//PositionRole 身份-角色
	PositionRole

	//UserOffice 用户-组织
	UserOffice
	//UserModule 用户-模块
	UserModule
	//UserRole 用户-角色
	UserRole
	//UserGroup 用户-组
	UserGroup
	//UserPosition 用户-身份
	UserPosition
)
