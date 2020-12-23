// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/11/6.

package module

import (
	"auth-gateway/serializer"
	"auth-gateway/service/module"
	"github.com/gin-gonic/gin"
)

// Create 创建模块
func Create(c *gin.Context) *serializer.Response {
	var rs module.CreateService
	if err := c.ShouldBind(&rs); err != nil {
		return serializer.I18Error(err)
	}
	if e := rs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Delete 删除模块
func Delete(c *gin.Context) *serializer.Response {
	var rs module.DeleteService
	if err := c.ShouldBind(&rs); err != nil {
		return serializer.I18Error(err)
	}
	if e := rs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// View 查看模块
func View(c *gin.Context) *serializer.Response {
	var rs module.ViewService
	if err := c.ShouldBind(&rs); err != nil {
		return serializer.I18Error(err)
	}
	v, e := rs.Execute()
	if nil != e {
		return serializer.Failed(e)
	}
	return serializer.Success(v)
}
