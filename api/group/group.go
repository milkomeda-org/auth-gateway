// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package group

import (
	"github.com/gin-gonic/gin"
	"oa-auth/serializer"
	"oa-auth/service/group"
)

// Create 创建用户组
func Create(c *gin.Context) serializer.Response {
	var gs group.GroupCreateService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Update 修改用户组
func Update(c *gin.Context) serializer.Response {
	var gs group.GroupUpdateService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Delete 删除用户组
func Delete(c *gin.Context) serializer.Response {
	var gs group.GroupDeleteService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// View 查看用户组
func View(c *gin.Context) serializer.Response {
	var gs group.GroupViewService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	v, e := gs.Execute()
	if e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(v)
}
