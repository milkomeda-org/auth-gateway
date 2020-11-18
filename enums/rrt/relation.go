// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package rrt

// ResRelationType 资源关联类型
type ResRelationType int

const (
	_ ResRelationType = iota

	GroupModule
	GroupRole

	PositionModule
	PositionRole

	UserOffice
	UserRole
	UserGroup
	UserPosition

	RoleModule
	RoleProxy
)
