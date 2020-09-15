// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package organization

import (
	"goa/serializer"
	"goa/service/organization"

	"github.com/gin-gonic/gin"
)

// GroupCreate 创建用户组
func GroupCreate(c *gin.Context) serializer.Response {
	var gs organization.GroupAddService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.ParamErr("", err)

	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// GroupUpdate 修改用户组
func GroupUpdate(c *gin.Context) serializer.Response {
	var gs organization.GroupUpdateService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.ParamErr("", err)
	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// GroupDelete 删除用户组
func GroupDelete(c *gin.Context) serializer.Response {
	var gs organization.GroupDeleteService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.ParamErr("", err)

	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
