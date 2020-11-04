// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"oa-auth/serializer"
	"oa-auth/service/organization"

	"github.com/gin-gonic/gin"
)

// PositionCreate 创建职位
func PositionCreate(c *gin.Context) serializer.Response {
	var ps organization.PositionCreateService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionUpdate 修改职位
func PositionUpdate(c *gin.Context) serializer.Response {
	var ps organization.PositionUpdateService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionDelete 删除职位
func PositionDelete(c *gin.Context) serializer.Response {
	var ps organization.PositionDeleteService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionView 查看职位
func PositionView(c *gin.Context) serializer.Response {
	var ps organization.PositionViewService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)
	}
	v, e := ps.Execute()
	if nil != e {
		return serializer.Failed(e)
	}
	return serializer.Success(v)
}

// PositionRoleAdd 添加职位角色
func PositionRoleAdd(c *gin.Context) serializer.Response {
	var ps organization.PositionRoleMappingAddService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionRoleRemove 移除职位角色
func PositionRoleRemove(c *gin.Context) serializer.Response {
	var ps organization.PositionRoleMappingRemoveService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.ParamErr("", err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
