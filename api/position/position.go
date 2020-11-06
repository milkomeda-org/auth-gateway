// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/9/14.

package position

import (
	"oa-auth/serializer"
	"oa-auth/service/position"

	"github.com/gin-gonic/gin"
)

// Create 创建职位
func Create(c *gin.Context) serializer.Response {
	var ps position.CreateService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.I18Error(err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Update 修改职位
func Update(c *gin.Context) serializer.Response {
	var ps position.UpdateService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.I18Error(err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Delete 删除职位
func Delete(c *gin.Context) serializer.Response {
	var ps position.DeleteService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.I18Error(err)
	}
	if e := ps.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// View 查看职位
func View(c *gin.Context) serializer.Response {
	var ps position.ViewService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.I18Error(err)
	}
	v, e := ps.Execute()
	if nil != e {
		return serializer.Failed(e)
	}
	return serializer.Success(v)
}

// RoleAdd 添加职位角色
func RoleAdd(c *gin.Context) serializer.Response {
	var ps position.RoleService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.I18Error(err)
	}
	if e := ps.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// RoleRemove 移除职位角色
func RoleRemove(c *gin.Context) serializer.Response {
	var ps position.RoleService
	if err := c.ShouldBind(&ps); err != nil {
		return serializer.I18Error(err)
	}
	if e := ps.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// ModuleAdd 添加模块
func ModuleAdd(c *gin.Context) serializer.Response {
	var os position.ModuleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// ModuleRemove 移除模块
func ModuleRemove(c *gin.Context) serializer.Response {
	var os position.ModuleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
