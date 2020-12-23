// Copyright The Milkomeda Org. All rights reserved.
// Created by vinson on 2020/9/14.

package group

import (
	"auth-gateway/serializer"
	"auth-gateway/service/group"
	"github.com/gin-gonic/gin"
)

// Create 创建用户组
func Create(c *gin.Context) *serializer.Response {
	var gs group.CreateService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Update 修改用户组
func Update(c *gin.Context) *serializer.Response {
	var gs group.UpdateService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Delete 删除用户组
func Delete(c *gin.Context) *serializer.Response {
	var gs group.DeleteService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	if e := gs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// View 查看用户组
func View(c *gin.Context) *serializer.Response {
	var gs group.ViewService
	if err := c.ShouldBind(&gs); err != nil {
		return serializer.I18Error(err)
	}
	v, e := gs.Execute()
	if e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(v)
}

// RoleAdd 添加角色
func RoleAdd(c *gin.Context) *serializer.Response {
	var os group.RoleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// RoleRemove 移除角色
func RoleRemove(c *gin.Context) *serializer.Response {
	var os group.RoleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// ModuleAdd 添加模块
func ModuleAdd(c *gin.Context) *serializer.Response {
	var os group.ModuleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// ModuleRemove 移除模块
func ModuleRemove(c *gin.Context) *serializer.Response {
	var os group.ModuleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
